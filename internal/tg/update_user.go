package tg

import (
	"fmt"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) UpdateUserInChatFromApi(username, phone *string, title string) error {
	err := t.updateUserFlow(title, username, phone)
	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to update user, some strange error")
			return errors.Wrap(err, "failed to update user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.UpdateUserInChatFromApi(username, phone, title)
		}

		t.log.WithError(err).Errorf("failed to kick user from %s", title)
		return errors.Wrap(err, fmt.Sprintf("failed to kick user from %s", title))
	}

	t.log.Infof("successfully update user in %s", title)
	return nil
}

func (t *tg) updateUserFlow(title string, username, phone *string) error {
	id, accessHash, err := t.findChatByTitle(title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to find chat %s", title)
		return err
	}

	chatUser, err := t.getChatUserFlow(username, phone, title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get chat user")
		return err
	}
	if chatUser == nil {
		t.log.Errorf("user is not in `%s`", title)
		return errors.Errorf("user is not in `%s`", title)
	}

	inputUser, err := t.getInputUser(username, phone)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get input user")
		return err
	}

	switch chatUser.AccessLevel {
	case data.Admin:
		if err = t.updateAdminToMember(inputUser, *id, accessHash); err != nil {
			t.log.WithError(err).Errorf("failed to update admin to member")
			return err
		}
	case data.Member:
		if err = t.updateMemberToAdmin(inputUser, *id, accessHash); err != nil {
			t.log.WithError(err).Errorf("failed to update member to admin")
			return err
		}
	case data.Owner:
		t.log.WithError(err).Errorf("can't update owner")
		return errors.New("can't update owner")
	case data.Left:
		t.log.WithError(err).Errorf("can't update user that left chat")
		return errors.New("can't update user that left chat")
	case data.Self:
		t.log.WithError(err).Errorf("can't update self")
		return errors.New("can't update self")
	case data.Banned:
		t.log.WithError(err).Errorf("can't update banned user")
		return errors.New("can't update banned user")
	default:
		t.log.Errorf("unexpected user status")
		return errors.Errorf("unexpected user status")
	}

	return nil
}

func (t *tg) updateMemberToAdmin(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsEditAdmin(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, user, &telegram.ChatAdminRights{
			ChangeInfo:     true,
			PostMessages:   true,
			EditMessages:   true,
			DeleteMessages: true,
			BanUsers:       true,
			InviteUsers:    true,
			PinMessages:    true,
			AddAdmins:      false,
			Anonymous:      false,
			ManageCall:     true,
		}, "")
		if err != nil {
			t.log.WithError(err).Errorf("failed to make user admin in channel")
			return err
		}
	} else {
		_, err := t.client.MessagesEditChatAdmin(id, user, true)
		if err != nil {
			t.log.WithError(err).Errorf("failed to make user admin in chat")
			return err
		}
	}

	return nil
}

func (t *tg) updateAdminToMember(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsEditAdmin(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, user, &telegram.ChatAdminRights{
			ChangeInfo:     false,
			PostMessages:   false,
			EditMessages:   false,
			DeleteMessages: false,
			BanUsers:       false,
			InviteUsers:    false,
			PinMessages:    false,
			AddAdmins:      false,
			Anonymous:      false,
			ManageCall:     false,
		}, "")
		if err != nil {
			t.log.WithError(err).Errorf("failed to make user member in channel")
			return err
		}
	} else {
		_, err := t.client.MessagesEditChatAdmin(id, user, false)
		if err != nil {
			t.log.WithError(err).Errorf("failed to make user member in chat")
			return err
		}
	}

	return nil
}

func (t *tg) getInputUser(username, phone *string) (*telegram.InputUserObj, error) {
	var inputUser *telegram.InputUserObj
	var err error
	if username != nil {
		inputUser, err = t.getUserByUsername(*username)
	}
	if phone != nil {
		inputUser, err = t.getUserByPhone(*phone)
	}
	if err != nil {
		t.log.WithError(err).Errorf("failed to get user")
		return nil, err
	}

	return inputUser, nil
}

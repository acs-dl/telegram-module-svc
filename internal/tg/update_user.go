package tg

import (
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
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

	var inputUser *telegram.InputUserObj
	if username != nil {
		inputUser, err = t.getUserByUsername(*username)
	}
	if phone != nil {
		inputUser, err = t.getUserByPhone(*phone)
	}
	if err != nil {
		t.log.WithError(err).Errorf("failed to get user")
		return err
	}

	user, err := t.checkUserInChat(*id, accessHash, int64(inputUser.UserID))
	if err != nil {
		t.log.WithError(err).Errorf("failed to check chat and user")
		return err
	}
	if user == nil {
		t.log.Errorf("user is not in `%s`", title)
		return errors.Errorf("user is not in `%s`", title)
	}

	status, err := t.getUserStatus(*id, accessHash, int64(inputUser.UserID))
	if err != nil {
		t.log.WithError(err).Errorf("failed to get user status")
		return err
	}

	//TODO: decide how to update from `admin` to `member` and from `member` to `admin`
	switch status {
	case "admin":
	case "owner":
		t.log.WithError(err).Errorf("can't update owner")
		return errors.New("can't update owner")
	case "member":
	default:
		t.log.Errorf("unexpected user status")
		return errors.Errorf("unexpected user status")
	}

	if err = t.updateUser(inputUser, *id, accessHash); err != nil {
		t.log.WithError(err).Errorf("failed to kick")
		return err
	}

	return nil
}

func (t *tg) updateUser(user *telegram.InputUserObj, id int32, hashID *int64) error {
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

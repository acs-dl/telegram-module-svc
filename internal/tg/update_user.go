package tg

import (
	"syscall"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) UpdateUserInChatFromApi(chatUser data.User, chat Chat) error {
	err := t.updateUserFlow(chatUser, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.UpdateUserInChatFromApi(chatUser, chat)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to update user, some strange error")
			return errors.Wrap(err, "failed to update user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.UpdateUserInChatFromApi(chatUser, chat)
		}

		t.log.WithError(err).Errorf("failed to update user in chat")
		return errors.Wrap(err, "failed to update user in chat")
	}

	t.log.Infof("successfully update user in chat")
	return nil
}

func (t *tg) updateUserFlow(chatUser data.User, chat Chat) error {
	switch chatUser.AccessLevel {
	case data.Admin:
		if err := t.updateAdminToMember(&telegram.InputUserObj{
			UserID:     int32(chatUser.TelegramId),
			AccessHash: chatUser.AccessHash,
		}, chat.id, chat.accessHash); err != nil {
			t.log.Errorf("failed to update admin to member")
			return err
		}
	case data.Member:
		if err := t.updateMemberToAdmin(&telegram.InputUserObj{
			UserID:     int32(chatUser.TelegramId),
			AccessHash: chatUser.AccessHash,
		}, chat.id, chat.accessHash); err != nil {
			t.log.Errorf("failed to update member to admin")
			return err
		}
	case data.Owner:
		t.log.Errorf("can't update owner")
		return errors.New("can't update owner")
	case data.Left:
		t.log.Errorf("can't update user that left chat")
		return errors.New("can't update user that left chat")
	case data.Self:
		t.log.Errorf("can't update self")
		return errors.New("can't update self")
	case data.Banned:
		t.log.Errorf("can't update banned user")
		return errors.New("can't update banned user")
	default:
		t.log.Errorf("unexpected user status")
		return errors.New("unexpected user status")
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
			t.log.Errorf("failed to make user admin in channel")
			return err
		}
	} else {
		_, err := t.client.MessagesEditChatAdmin(id, user, true)
		if err != nil {
			t.log.Errorf("failed to make user admin in chat")
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
			t.log.Errorf("failed to make user member in channel")
			return err
		}
	} else {
		_, err := t.client.MessagesEditChatAdmin(id, user, false)
		if err != nil {
			t.log.Errorf("failed to make user member in chat")
			return err
		}
	}

	return nil
}

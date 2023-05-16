package tg_client

import (
	"syscall"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) UpdateUserInChatFromApi(chatUser data.User, chat Chat) error {
	err := t.updateUserFlow(chatUser, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.UpdateUserInChatFromApi(chatUser, chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.UpdateUserInChatFromApi(chatUser, chat)
		}

		t.log.WithError(err).Errorf("failed to update user in chat")
		return errors.Wrap(err, "failed to update user in chat")
	}

	t.log.Infof("successfully update user in chat")
	return nil
}

func (t *tgInfo) updateUserFlow(chatUser data.User, chat Chat) error {
	switch chatUser.AccessLevel {
	case data.Admin:
		if err := t.updateAdminToMember(&tg.InputUser{
			UserID:     chatUser.TelegramId,
			AccessHash: chatUser.AccessHash,
		}, chat.Id, chat.AccessHash); err != nil {
			t.log.Errorf("failed to update admin to member")
			return err
		}
	case data.Member:
		if err := t.updateMemberToAdmin(&tg.InputUser{
			UserID:     chatUser.TelegramId,
			AccessHash: chatUser.AccessHash,
		}, chat.Id, chat.AccessHash); err != nil {
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

func (t *tgInfo) updateMemberToAdmin(user *tg.InputUser, id int64, hashID *int64) error {
	if hashID != nil {
		_, err := t.superUserClient.API().ChannelsEditAdmin(t.ctx, &tg.ChannelsEditAdminRequest{
			Channel: &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
			UserID:  user,
			AdminRights: tg.ChatAdminRights{
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
			},
			Rank: ""})
		if err != nil {
			t.log.Errorf("failed to make user admin in channel")
			return err
		}
	} else {
		_, err := t.superUserClient.API().MessagesEditChatAdmin(t.ctx, &tg.MessagesEditChatAdminRequest{
			ChatID:  id,
			UserID:  user,
			IsAdmin: true,
		})
		if err != nil {
			t.log.Errorf("failed to make user admin in chat")
			return err
		}
	}

	return nil
}

func (t *tgInfo) updateAdminToMember(user *tg.InputUser, id int64, hashID *int64) error {
	if hashID != nil {
		_, err := t.superUserClient.API().ChannelsEditAdmin(t.ctx, &tg.ChannelsEditAdminRequest{
			Channel: &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
			UserID:  user,
			AdminRights: tg.ChatAdminRights{
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
			},
			Rank: ""})
		if err != nil {
			t.log.Errorf("failed to make user member in channel")
			return err
		}
	} else {
		_, err := t.superUserClient.API().MessagesEditChatAdmin(t.ctx, &tg.MessagesEditChatAdminRequest{
			ChatID:  id,
			UserID:  user,
			IsAdmin: false,
		})
		if err != nil {
			t.log.Errorf("failed to make user member in chat")
			return err
		}
	}

	return nil
}

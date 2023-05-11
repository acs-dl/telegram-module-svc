package tg_client

import (
	"syscall"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
)

func (t *tgInfo) AddUserInChatFromApi(user data.User, chat Chat) error {
	err := t.addUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.AddUserInChatFromApi(user, chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.AddUserInChatFromApi(user, chat)
		}

		return err
	}

	t.log.Infof("successfully add user in chat")
	return nil
}

func (t *tgInfo) addUserFlow(user data.User, chat Chat) error {
	if err := t.addUser(&tg.InputUser{
		UserID:     user.TelegramId,
		AccessHash: user.AccessHash,
	}, chat.id, chat.accessHash); err != nil {
		return err
	}

	return nil
}

func (t *tgInfo) addUser(user *tg.InputUser, id int64, hashID *int64) error {
	if hashID != nil {
		_, err := t.superUserClient.API().ChannelsInviteToChannel(t.ctx, &tg.ChannelsInviteToChannelRequest{
			Channel: &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
			Users:   []tg.InputUserClass{user}})
		if err != nil {
			t.log.Errorf("failed to invite to channel")
			return err
		}
	} else {
		_, err := t.superUserClient.API().MessagesAddChatUser(t.ctx, &tg.MessagesAddChatUserRequest{
			ChatID:   id,
			UserID:   user,
			FwdLimit: 0,
		})
		if err != nil {
			t.log.Errorf("failed to add chat user")
			return err
		}
	}

	return nil
}

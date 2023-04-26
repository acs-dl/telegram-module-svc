package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) AddUserInChatFromApi(user data.User, chat Chat) error {
	err := t.addUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superClient = cl.GetTg().superClient
			return t.AddUserInChatFromApi(user, chat)
		}

		if tgerr.IsCode(err, 420) {
			duration, ok := tgerr.AsFloodWait(err)
			if !ok {
				return errors.New("failed to convert flood error")
			}
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.AddUserInChatFromApi(user, chat)
		}

		return errors.Wrap(err, "failed to add user")
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
		_, err := t.superClient.API().ChannelsInviteToChannel(t.ctx, &tg.ChannelsInviteToChannelRequest{
			Channel: &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
			Users:   []tg.InputUserClass{user}})
		if err != nil {
			t.log.Errorf("failed to invite to channel")
			return err
		}
	} else {
		_, err := t.superClient.API().MessagesAddChatUser(t.ctx, &tg.MessagesAddChatUserRequest{
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

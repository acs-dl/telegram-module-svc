package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetChatUserFromApi(user data.User, chat Chat) (*data.User, error) {
	chatUser, err := t.getChatUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superClient = cl.GetTg().superClient
			return t.GetChatUserFromApi(user, chat)
		}

		if tgerr.IsCode(err, 420) {
			duration, ok := tgerr.AsFloodWait(err)
			if !ok {
				return nil, errors.New("failed to convert flood error")
			}
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetChatUserFromApi(user, chat)
		}

		t.log.WithError(err).Errorf("failed to get chat user")
		return nil, errors.Wrap(err, "failed to get chat user")
	}

	t.log.Infof("successfully got chat user")
	return chatUser, nil
}

func (t *tgInfo) getChatUserFlow(user data.User, chat Chat) (*data.User, error) {
	var chatUsers, err = t.getAllUsers(chat.id, chat.accessHash)
	if err != nil {
		t.log.Errorf("failed to get all users")
		return nil, err
	}

	for _, chatUser := range chatUsers {
		if user.TelegramId == chatUser.TelegramId {
			return &chatUser, nil
		}
	}

	return nil, nil
}

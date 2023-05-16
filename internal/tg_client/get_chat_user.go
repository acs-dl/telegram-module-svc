package tg_client

import (
	"syscall"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetChatUserFromApi(user data.User, chat Chat) (*data.User, error) {
	chatUser, err := t.getChatUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.GetChatUserFromApi(user, chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
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
	var chatUsers, err = t.getAllUsers(chat.Id, chat.AccessHash)
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

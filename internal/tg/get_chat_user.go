package tg

import (
	"syscall"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) GetChatUserFromApi(user data.User, chat Chat) (*data.User, error) {
	chatUser, err := t.getChatUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.GetChatUserFromApi(user, chat)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to get chat user, some strange error")
			return nil, errors.Wrap(err, "failed to get chat user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.GetChatUserFromApi(user, chat)
		}

		t.log.WithError(err).Errorf("failed to get chat user")
		return nil, errors.Wrap(err, "failed to get chat user")
	}

	t.log.Infof("successfully got chat user")
	return chatUser, nil
}

func (t *tg) getChatUserFlow(user data.User, chat Chat) (*data.User, error) {
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

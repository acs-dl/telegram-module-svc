package tg

import (
	"fmt"
	"syscall"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) GetChatUserFromApi(username, phone *string, title string) (*data.User, error) {
	user, err := t.getChatUserFlow(username, phone, title)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.GetChatUserFromApi(username, phone, title)
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
			return t.GetChatUserFromApi(username, phone, title)
		}

		t.log.WithError(err).Errorf("failed to get chat user")
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get chat user"))
	}

	t.log.Infof("successfully got chat user")
	return user, nil
}

func (t *tg) getChatUserFlow(username, phone *string, title string) (*data.User, error) {
	id, accessHash, err := t.GetChatFromApi(title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to find chat %s", title)
		return nil, err
	}

	user, err := t.GetUserFromApi(username, phone)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get user")
		return nil, err
	}

	return t.checkUserInChat(*id, accessHash, user.TelegramId)
}

func (t *tg) checkUserInChat(id int32, hashID *int64, userId int64) (*data.User, error) {
	var users, err = t.getAllUsers(id, hashID)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get all users")
		return nil, err
	}

	for _, user := range users {
		if userId == user.TelegramId {
			return &user, nil
		}
	}

	return nil, nil
}

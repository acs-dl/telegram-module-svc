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

func (t *tg) AddUserInChatFromApi(user data.User, chat Chat) error {
	err := t.addUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.AddUserInChatFromApi(user, chat)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to add user, some strange error")
			return errors.Wrap(err, "failed to add user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.AddUserInChatFromApi(user, chat)
		}

		return errors.Wrap(err, "failed to add user")
	}

	t.log.Infof("successfully add user in chat")
	return nil
}

func (t *tg) addUserFlow(user data.User, chat Chat) error {
	if err := t.addUser(&telegram.InputUserObj{
		UserID:     int32(user.TelegramId),
		AccessHash: user.AccessHash,
	}, chat.id, chat.accessHash); err != nil {
		return err
	}

	return nil
}

func (t *tg) addUser(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsInviteToChannel(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, []telegram.InputUser{user})
		if err != nil {
			t.log.Errorf("failed to invite to channel")
			return err
		}
	} else {
		_, err := t.client.MessagesAddChatUser(id, user, 0)
		if err != nil {
			t.log.Errorf("failed to add chat user")
			return err
		}
	}

	return nil
}

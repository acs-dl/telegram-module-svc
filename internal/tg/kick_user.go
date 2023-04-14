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

func (t *tg) DeleteFromChatFromApi(user data.User, chat Chat) error {
	err := t.kickUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.DeleteFromChatFromApi(user, chat)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to kick user, some strange error")
			return errors.Wrap(err, "failed to kick user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.DeleteFromChatFromApi(user, chat)
		}

		t.log.WithError(err).Errorf("failed to kick user")
		return errors.Wrap(err, "failed to kick user")
	}

	t.log.Infof("successfully kicked user")
	return nil
}

func (t *tg) kickUserFlow(user data.User, chat Chat) error {
	if err := t.kickUser(&telegram.InputUserObj{
		UserID:     int32(user.TelegramId),
		AccessHash: user.AccessHash,
	}, chat.id, chat.accessHash); err != nil {
		t.log.Errorf("failed to kick user")
		return err
	}

	return nil
}

func (t *tg) kickUser(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsEditBanned(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, user, &telegram.ChatBannedRights{
			ViewMessages: true,
			SendMessages: true,
			SendMedia:    true,
			SendStickers: true,
			SendGifs:     true,
			SendGames:    true,
			SendInline:   true,
			EmbedLinks:   true,
			SendPolls:    true,
			ChangeInfo:   true,
			InviteUsers:  true,
			PinMessages:  true,
			UntilDate:    0,
		})
		if err != nil {
			t.log.Errorf("failed to ban channel user")
			return err
		}
	} else {
		_, err := t.client.MessagesDeleteChatUser(id, user)
		if err != nil {
			t.log.Errorf("failed to delete chat user")
			return err
		}
	}

	return nil
}

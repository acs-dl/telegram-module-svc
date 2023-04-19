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

func (t *tgInfo) DeleteFromChatFromApi(user data.User, chat Chat) error {
	err := t.kickUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superClient = cl.GetTg().superClient
			return t.DeleteFromChatFromApi(user, chat)
		}

		if tgerr.IsCode(err, 420) {
			duration, ok := tgerr.AsFloodWait(err)
			if !ok {
				return errors.New("failed to convert flood error")
			}
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.DeleteFromChatFromApi(user, chat)
		}

		t.log.WithError(err).Errorf("failed to kick user")
		return errors.Wrap(err, "failed to kick user")
	}

	t.log.Infof("successfully kicked user")
	return nil
}

func (t *tgInfo) kickUserFlow(user data.User, chat Chat) error {
	if err := t.kickUser(&tg.InputUser{
		UserID:     user.TelegramId,
		AccessHash: user.AccessHash,
	}, chat.id, chat.accessHash); err != nil {
		t.log.Errorf("failed to kick user")
		return err
	}

	return nil
}

func (t *tgInfo) kickUser(user *tg.InputUser, id int64, hashID *int64) error {
	if hashID != nil {
		_, err := t.superClient.API().ChannelsEditBanned(t.ctx, &tg.ChannelsEditBannedRequest{
			Channel:     &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
			Participant: &tg.InputPeerUser{UserID: user.UserID, AccessHash: user.AccessHash},
			BannedRights: tg.ChatBannedRights{
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
			}})
		if err != nil {
			t.log.Errorf("failed to ban channel user")
			return err
		}
	} else {
		_, err := t.superClient.API().MessagesDeleteChatUser(t.ctx, &tg.MessagesDeleteChatUserRequest{ChatID: id, UserID: user})
		if err != nil {
			t.log.Errorf("failed to delete chat user")
			return err
		}
	}

	return nil
}

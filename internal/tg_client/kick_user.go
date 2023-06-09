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

func (t *tgInfo) DeleteFromChatFromApi(user data.User, chat Chat) error {
	err := t.kickUserFlow(user, chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.DeleteFromChatFromApi(user, chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
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
	}, chat.Id, chat.AccessHash); err != nil {
		t.log.Errorf("failed to kick user")
		return err
	}

	return nil
}

func (t *tgInfo) kickUser(user *tg.InputUser, id int64, hashID *int64) error {
	if hashID != nil {
		_, err := t.superUserClient.API().ChannelsEditBanned(t.ctx, &tg.ChannelsEditBannedRequest{
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
				UntilDate:    int(time.Now().Add(40 * time.Second).Unix()),
			}})
		if err != nil {
			t.log.Errorf("failed to ban channel user")
			return err
		}
	} else {
		_, err := t.superUserClient.API().MessagesDeleteChatUser(t.ctx, &tg.MessagesDeleteChatUserRequest{ChatID: id, UserID: user})
		if err != nil {
			t.log.Errorf("failed to delete chat user")
			return err
		}
	}

	return nil
}

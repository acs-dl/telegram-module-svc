package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GenerateChatLinkFromApi(chat Chat) (string, error) {
	inviteLink, err := t.generateChatLink(chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.GenerateChatLinkFromApi(chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GenerateChatLinkFromApi(chat)
		}

		return "", errors.Wrap(err, "failed to generate invite link")
	}

	t.log.Infof("successfully generate invite link")
	return inviteLink, nil
}

func (t *tgInfo) generateChatLink(chat Chat) (string, error) {
	inputPeer := tg.InputPeerClass(nil)

	if chat.AccessHash != nil {
		inputPeer = &tg.InputPeerChannel{ChannelID: chat.Id, AccessHash: *chat.AccessHash}
	} else {
		inputPeer = &tg.InputPeerChat{ChatID: chat.Id}
	}

	expireDate := time.Now().Add(1 * time.Hour).Unix()
	inviteClass, err := t.superUserClient.API().MessagesExportChatInvite(t.ctx, &tg.MessagesExportChatInviteRequest{
		Peer:       inputPeer,
		UsageLimit: 1,
		ExpireDate: int(expireDate),
	})
	if err != nil {
		return "", err
	}

	invite, ok := inviteClass.(*tg.ChatInviteExported)
	if !ok {
		return "", errors.New("wrong invite type")
	}

	return invite.Link, nil
}

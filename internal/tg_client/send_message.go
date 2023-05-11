package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/html"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) SendMessageFromApi(info data.MessageInfo) error {
	err := t.sendMessage(info)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.SendMessageFromApi(info)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.SendMessageFromApi(info)
		}

		return errors.Wrap(err, "failed to send message")
	}

	t.log.Infof("successfully send message")
	return nil
}

func (t *tgInfo) sendMessage(info data.MessageInfo) error {
	msg, err := GenerateStringFromTemplate(info.Data, info.MessageTemplate)
	if err != nil {
		return err
	}

	messageSender := message.NewSender(t.superUserClient.API())

	_, err = messageSender.To(&tg.InputPeerUser{
		UserID:     info.User.TelegramId,
		AccessHash: info.User.AccessHash,
	}).StyledText(t.ctx, html.String(nil, msg))
	if err != nil {
		return err
	}

	return nil
}

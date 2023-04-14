package tg

import (
	"fmt"
	"syscall"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) GetChatFromApi(title string) (*Chat, error) {
	chat, err := t.getChatFlow(title)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.GetChatFromApi(title)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to get chat, some strange error")
			return nil, errors.Wrap(err, "failed to get chat, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.GetChatFromApi(title)
		}

		return nil, errors.Wrap(err, fmt.Sprintf("failed to get chat `%s`", title))
	}

	t.log.Infof("successfully got chat")
	return chat, nil
}

func (t *tg) getChatFlow(title string) (*Chat, error) {
	chat, err := t.findChatByTitle(title)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (t *tg) findChatByTitle(title string) (*Chat, error) {
	discussion, err := t.client.MessagesGetAllChats(nil)
	if err != nil {
		return nil, err
	}

	for _, chat := range discussion.(*telegram.MessagesChatsObj).Chats {
		switch converted := chat.(type) {
		case *telegram.Channel:
			if converted.Title == title {
				return &Chat{
					converted.ID,
					&converted.AccessHash,
				}, nil
			}
		case *telegram.ChatObj:
			if converted.MigratedTo != nil {
				continue //it means that chat migrated to channel (supergroup)
			}
			if converted.Title == title {
				return &Chat{
					converted.ID,
					nil,
				}, nil
			}
		default:
			return nil, errors.Errorf("unexpected type %T", converted)
		}
	}

	t.log.Errorf("no chat `%s` was found", title)
	return nil, nil
}

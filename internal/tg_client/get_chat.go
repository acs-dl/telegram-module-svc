package tg_client

import (
	"fmt"
	"syscall"
	"time"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetChatFromApi(title string) (*Chat, error) {
	chat, err := t.getChatFlow(title)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superClient = cl.GetTg().superClient
			return t.GetChatFromApi(title)
		}

		if tgerr.IsCode(err, 420) {
			duration, ok := tgerr.AsFloodWait(err)
			if !ok {
				return nil, errors.New("failed to convert flood error")
			}
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetChatFromApi(title)
		}

		return nil, errors.Wrap(err, fmt.Sprintf("failed to get chat `%s`", title))
	}

	t.log.Infof("successfully got chat")
	return chat, nil
}

func (t *tgInfo) getChatFlow(title string) (*Chat, error) {
	chat, err := t.findChatByTitle(title)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (t *tgInfo) findChatByTitle(title string) (*Chat, error) {
	searched, err := t.superClient.API().ContactsSearch(t.ctx, &tg.ContactsSearchRequest{
		Q:     title,
		Limit: 10,
	})
	//discussion, err := t.superClient.API().MessagesGetAllChats(t.ctx, nil)
	if err != nil {
		return nil, err
	}

	//for _, chat := range discussion.(*tg.MessagesChats).Chats {
	for _, chat := range searched.Chats {
		switch converted := chat.(type) {
		case *tg.Channel:
			if converted.Title == title {
				return &Chat{
					converted.ID,
					&converted.AccessHash,
				}, nil
			}
		case *tg.Chat:
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

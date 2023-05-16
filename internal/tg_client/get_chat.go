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

func (t *tgInfo) GetChatFromApi(title string) ([]Chat, error) {
	chats, err := t.getChatFlow(title)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.GetChatFromApi(title)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetChatFromApi(title)
		}

		return nil, errors.Wrap(err, fmt.Sprintf("failed to get chat `%s`", title))
	}

	t.log.Infof("successfully got chat")
	return chats, nil
}

func (t *tgInfo) getChatFlow(title string) ([]Chat, error) {
	chats, err := t.findChatByTitle(title)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (t *tgInfo) findChatByTitle(title string) ([]Chat, error) {
	searched, err := t.superUserClient.API().ContactsSearch(t.ctx, &tg.ContactsSearchRequest{
		Q:     title,
		Limit: 10,
	})
	if err != nil {
		return nil, err
	}

	chats := make([]Chat, 0)

	for _, chat := range searched.Chats {
		switch converted := chat.(type) {
		case *tg.Channel:
			if converted.Title == title {
				chats = append(chats, Chat{
					Title:         converted.Title,
					Id:            converted.ID,
					AccessHash:    &converted.AccessHash,
					Photo:         &converted.Photo,
					MembersAmount: int64(converted.ParticipantsCount),
				})
			}
		case *tg.Chat:
			if converted.MigratedTo != nil {
				continue //it means that chat migrated to channel (supergroup)
			}
			if converted.Title == title {
				chats = append(chats, Chat{
					Title:         converted.Title,
					Id:            converted.ID,
					AccessHash:    nil,
					Photo:         &converted.Photo,
					MembersAmount: int64(converted.ParticipantsCount),
				})
			}
		default:
			return nil, errors.Errorf("unexpected type %T", converted)
		}
	}

	return chats, nil
}

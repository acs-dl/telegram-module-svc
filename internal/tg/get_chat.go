package tg

import (
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) findChatByTitle(title string) (*int32, *int64, error) {
	discussion, err := t.client.MessagesGetAllChats(nil)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get all chats")
		return nil, nil, err
	}

	for _, chat := range discussion.(*telegram.MessagesChatsObj).Chats {
		switch converted := chat.(type) {
		default:
			t.log.Errorf("unexpected chat type %T", converted)
			return nil, nil, errors.Errorf("unexpected type %T", converted)
		case *telegram.Channel:
			if converted.Title == title {
				return &converted.ID, &converted.AccessHash, nil
			}
		case *telegram.ChatObj:
			if converted.MigratedTo != nil {
				continue //it means that chat migrated to channel (supergroup)
			}
			if converted.Title == title {
				return &converted.ID, nil, nil
			}
		}
	}

	t.log.Errorf("no chat `%s` was found", title)
	return nil, nil, errors.Errorf("no chat `%s` was found", title)
}

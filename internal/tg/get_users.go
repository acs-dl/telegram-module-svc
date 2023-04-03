package tg

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/google/uuid"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) GetUsersFromApi(title string) ([]data.User, error) {
	myInterface1 := interface{}(t.GetUsersFromApi)
	myInterface2 := interface{}(t.UpdateUserInChatFromApi)
	myInterface3 := interface{}(t.DeleteFromChatFromApi)
	//fmt.Println(reflect.TypeOf(myInterface3))
	//pqueue.TestHeap(map[interface{}]int{myInterface1: 1, myInterface2: 5, myInterface3: 3})

	pq := make(pqueue.PriorityQueue, 0)

	item1 := &pqueue.QueueItem{
		Uuid:     uuid.New(),
		Func:     myInterface3,
		Priority: 10,
	}
	heap.Push(&pq, item1)

	item3 := &pqueue.QueueItem{
		Uuid:     uuid.New(),
		Func:     myInterface1,
		Priority: 5,
	}
	heap.Push(&pq, item3)

	item2 := &pqueue.QueueItem{
		Uuid:     uuid.New(),
		Func:     myInterface2,
		Priority: 15,
	}
	heap.Push(&pq, item2)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqueue.QueueItem)
		fmt.Printf("%s, %d, %s \n", item.Uuid.String(), item.Priority, item.Func)
	}

	/* OUTPUT:
	79602114-5227-49bc-aec7-7f2ca322bf06, 15, %!s(func(*string, *string, string) error=0x9e64c0)
	01343912-33cf-43b5-bc57-2e720bacd1a4, 10, %!s(func(*string, *string, string) error=0x9e6560)
	79b1c46d-b423-4bf2-8968-fbe93b42d063, 5, %!s(func(string) ([]data.User, error)=0x9e6440)
	*/

	panic("hard stop")
	return nil, nil
	//users, err := t.getChatMembersByTitle(title)
	//if err != nil {
	//	if pkgErrors.Is(err, syscall.EPIPE) {
	//		cl := NewTg(t.tgCfg, t.log)
	//		t.client = cl.GetClient()
	//		return t.GetUsersFromApi(title)
	//	}
	//
	//	errResponse := &mtproto.ErrResponseCode{}
	//	if !pkgErrors.As(err, &errResponse) {
	//		return nil, errors.Wrap(err, "failed to get chat members, some strange error")
	//	}
	//	if errResponse.Message == "FLOOD_WAIT_X" {
	//		timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
	//		t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
	//		time.Sleep(timeoutDuration)
	//		return t.GetUsersFromApi(title)
	//	}
	//
	//	t.log.WithError(err).Errorf("failed to get chat members")
	//	return nil, errors.Wrap(err, "failed to get chat members")
	//}

	//return users, nil
}

func (t *tg) getChatMembersByTitle(title string) ([]data.User, error) {
	var users []data.User

	id, accessHash, err := t.findChatByTitle(title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to find chat %s", title)
		return nil, err
	}

	users, err = t.getAllUsers(*id, accessHash)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get all users")
		return nil, err
	}

	return users, nil
}

func (t *tg) getAllUsers(id int32, hashID *int64) ([]data.User, error) {
	var users []data.User

	if hashID != nil {
		channelUsers, err := t.getAllUsersFromChannel(id, hashID)
		if err != nil {
			t.log.WithError(err).Errorf("failed to get all users from channel")
			return nil, err
		}
		users = channelUsers
	} else {
		chatUsers, err := t.getAllUsersFromChat(id)
		if err != nil {
			t.log.WithError(err).Errorf("failed to get all users from chat")
			return nil, err
		}
		users = chatUsers
	}

	if len(users) == 0 {
		t.log.Errorf("no users in chat with id `%d` and hash `%d`", id, hashID)
		return nil, errors.Errorf("no users in chat with id `%d` and hash `%d`", id, hashID)
	}

	t.log.Infof("found `%d` users", len(users))
	return users, nil
}

func (t *tg) getAllUsersFromChat(chatId int32) ([]data.User, error) {
	fullChat, err := t.client.MessagesGetFullChat(chatId)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get full chat")
		return nil, err
	}

	userStatus := map[int32]string{}
	switch full := fullChat.FullChat.(type) {
	case *telegram.ChannelFull:
		//can't be so, because this `MessagesGetFullChat` returns chat, not channel, BUT I suppose we need to process such case
		t.log.Errorf("chat can't be channel")
		return nil, errors.Errorf("chat can't be channel")
	case *telegram.ChatFullObj:
		for _, participant := range full.Participants.(*telegram.ChatParticipantsObj).Participants {
			switch user := participant.(type) {
			case *telegram.ChatParticipantObj: // member
				userStatus[user.UserID] = data.Member
			case *telegram.ChatParticipantCreator: // owner
				userStatus[user.UserID] = data.Owner
			case *telegram.ChatParticipantAdmin: // admin
				userStatus[user.UserID] = data.Admin
			default:
				t.log.Errorf("unexpected user type %T", user)
				return nil, errors.Errorf("unexpected user type %T", user)
			}
		}
	default:
		t.log.Errorf("unexpected chat type %T", full)
		return nil, errors.Errorf("unexpected chat type %T", full)

	}

	var users []data.User

	for i := range fullChat.Users {
		tgUser := fullChat.Users[i].(*telegram.UserObj)
		users = append(users, data.User{
			Username:    &tgUser.Username,
			Phone:       &tgUser.Phone,
			FirstName:   tgUser.FirstName,
			LastName:    tgUser.LastName,
			TelegramId:  int64(tgUser.ID),
			AccessHash:  tgUser.AccessHash,
			AccessLevel: userStatus[tgUser.ID],
		})
	}

	return users, nil
}

func (t *tg) getAllUsersFromChannel(id int32, hashID *int64) ([]data.User, error) {
	var tgUsers []telegram.User
	var dbUsers []data.User
	var limit int32 = 1000
	var offset int32 = 0

	for {
		participants, err := t.client.ChannelsGetParticipants(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID}, &telegram.ChannelParticipantsSearch{}, offset, limit, 0)
		if err != nil {
			errResponse := &mtproto.ErrResponseCode{}
			if !pkgErrors.As(err, &errResponse) {
				t.log.WithError(err).Errorf("failed to get channel participants, some strange error")
				return nil, errors.Wrap(err, "failed to get channel participants, some strange error")
			}

			if errResponse.Message == "FLOOD_WAIT_X" {
				timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
				t.log.Warnf("we need to wait %s", timeoutDuration.String())
				time.Sleep(timeoutDuration)
				continue
			}

			t.log.WithError(err).Errorf("failed to get channel participants")
			return nil, errors.Wrap(err, "failed to get channel participants")
		}
		userStatus := map[int32]string{}
		participantsObject := participants.(*telegram.ChannelsChannelParticipantsObj)

		if len(participantsObject.Users) == 0 {
			break
		}

		for _, participant := range participantsObject.Participants {
			switch user := participant.(type) {
			case *telegram.ChannelParticipantSelf: // myself - it exists in tg api, when user (whose acc we use) isn't admin, but can get participant list user's status is 'self'
				userStatus[user.UserID] = data.Member
			case *telegram.ChannelParticipantObj: // member
				userStatus[user.UserID] = data.Member
			case *telegram.ChannelParticipantAdmin: // Admin
				userStatus[user.UserID] = data.Admin
			case *telegram.ChannelParticipantCreator: // Owner
				userStatus[user.UserID] = data.Owner
			case *telegram.ChannelParticipantBanned: // Banned/kicked user
				userStatus[user.UserID] = data.Banned
			case *telegram.ChannelParticipantLeft: // A participant that left the channel/supergroup
				userStatus[user.UserID] = data.Left
			default:
				t.log.Errorf("unexpected user type %T", user)
				return nil, errors.Errorf("unexpected user type %T", user)
			}
		}

		tgUsers = append(tgUsers, removeDuplicateUser(participantsObject.Users)...)

		for i := range participantsObject.Users {
			tgUser := participantsObject.Users[i].(*telegram.UserObj)
			if userStatus[tgUser.ID] == "" {
				continue
			}
			dbUsers = append(dbUsers, data.User{
				Username:    &tgUser.Username,
				Phone:       &tgUser.Phone,
				FirstName:   tgUser.FirstName,
				LastName:    tgUser.LastName,
				TelegramId:  int64(tgUser.ID),
				AccessHash:  tgUser.AccessHash,
				AccessLevel: userStatus[tgUser.ID],
			})
		}

		offset += limit
	}

	return dbUsers, nil
}

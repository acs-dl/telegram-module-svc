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

func (t *tg) GetChatUsersFromApi(chat Chat) ([]data.User, error) {
	users, err := t.getChatMembers(chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.GetChatUsersFromApi(chat)
		}

		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			return nil, errors.Wrap(err, "failed to get chat members, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.GetChatUsersFromApi(chat)
		}

		t.log.Errorf("failed to get chat members")
		return nil, errors.Wrap(err, "failed to get chat members")
	}

	return users, nil
}

func (t *tg) getChatMembers(chat Chat) ([]data.User, error) {
	users, err := t.getAllUsers(chat.id, chat.accessHash)
	if err != nil {
		t.log.Errorf("failed to get all users")
		return nil, err
	}

	return users, nil
}

func (t *tg) getAllUsers(id int32, hashID *int64) ([]data.User, error) {
	users := make([]data.User, 0)
	var err error = nil

	if hashID != nil {
		users, err = t.getAllUsersFromChannel(id, hashID)
		if err != nil {
			t.log.Errorf("failed to get all users from channel")
			return nil, err
		}
	} else {
		users, err = t.getAllUsersFromChat(id)
		if err != nil {
			t.log.Errorf("failed to get all users from chat")
			return nil, err
		}
	}

	t.log.Infof("found `%d` users", len(users))
	return users, nil
}

func (t *tg) getAllUsersFromChat(chatId int32) ([]data.User, error) {
	fullChat, err := t.client.MessagesGetFullChat(chatId)
	if err != nil {
		t.log.Errorf("failed to get full chat")
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
		//TODO: think how to handle several api req for getting users (if chat has more than 1000 users). Do we need PQueue just for it? :'(
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
			case *telegram.ChannelParticipantSelf: // myself - it exists in tg api, when user (whose acc we use) isn't admin, but can get participant list, user's status is 'self'
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

		//if we got all users with one api request, so we don't need to make more requests
		if int32(len(participantsObject.Users)) == participantsObject.Count {
			break
		}

		offset += limit
	}

	return dbUsers, nil
}

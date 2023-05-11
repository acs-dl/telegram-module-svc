package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetChatUsersFromApi(chat Chat) ([]data.User, error) {
	users, err := t.getChatMembers(chat)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.superUserClient = cl.GetTg().superUserClient
			return t.GetChatUsersFromApi(chat)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetChatUsersFromApi(chat)
		}

		t.log.Errorf("failed to get chat members")
		return nil, errors.Wrap(err, "failed to get chat members")
	}

	return users, nil
}

func (t *tgInfo) getChatMembers(chat Chat) ([]data.User, error) {
	users, err := t.getAllUsers(chat.id, chat.accessHash)
	if err != nil {
		t.log.Errorf("failed to get all users")
		return nil, err
	}

	return users, nil
}

func (t *tgInfo) getAllUsers(id int64, hashID *int64) ([]data.User, error) {
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

func (t *tgInfo) getAllUsersFromChat(chatId int64) ([]data.User, error) {
	fullChat, err := t.superUserClient.API().MessagesGetFullChat(t.ctx, chatId)
	if err != nil {
		t.log.Errorf("failed to get full chat")
		return nil, err
	}

	userStatus := map[int64]string{}
	switch full := fullChat.FullChat.(type) {
	case *tg.ChannelFull:
		//can't be so, because this `MessagesGetFullChat` returns chat, not channel, BUT I suppose we need to process such case
		t.log.Errorf("chat can't be channel")
		return nil, errors.Errorf("chat can't be channel")
	case *tg.ChatFull:
		for _, participant := range full.Participants.(*tg.ChatParticipants).Participants {
			switch user := participant.(type) {
			case *tg.ChatParticipant: // member
				userStatus[user.UserID] = data.Member
			case *tg.ChatParticipantCreator: // owner
				userStatus[user.UserID] = data.Owner
			case *tg.ChatParticipantAdmin: // admin
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
		tgUser := fullChat.Users[i].(*tg.User)
		users = append(users, data.User{
			Username:    &tgUser.Username,
			Phone:       &tgUser.Phone,
			FirstName:   tgUser.FirstName,
			LastName:    tgUser.LastName,
			TelegramId:  tgUser.ID,
			AccessHash:  tgUser.AccessHash,
			AccessLevel: userStatus[tgUser.ID],
		})
	}

	return users, nil
}

func (t *tgInfo) getAllUsersFromChannel(id int64, hashID *int64) ([]data.User, error) {
	var tgUsers []tg.User
	var dbUsers []data.User
	var limit = 1000
	var offset = 0

	for {
		//TODO: think how to handle several api req for getting users (if chat has more than 1000 users). Do we need PQueue just for it? :'(
		participants, err := t.superUserClient.API().ChannelsGetParticipants(t.ctx,
			&tg.ChannelsGetParticipantsRequest{
				Channel: &tg.InputChannel{ChannelID: id, AccessHash: *hashID},
				Filter:  &tg.ChannelParticipantsSearch{},
				Offset:  offset,
				Limit:   limit,
			},
		)
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
		userStatus := map[int64]string{}
		participantsObject := participants.(*tg.ChannelsChannelParticipants)

		if len(participantsObject.Users) == 0 {
			break
		}

		for _, participant := range participantsObject.Participants {
			switch user := participant.(type) {
			case *tg.ChannelParticipantSelf: // myself - it exists in tg_client api, when user (whose acc we use) isn't admin, but can get participant list, user's status is 'self'
				userStatus[user.UserID] = data.Member
			case *tg.ChannelParticipant: // member
				userStatus[user.UserID] = data.Member
			case *tg.ChannelParticipantAdmin: // Admin
				userStatus[user.UserID] = data.Admin
			case *tg.ChannelParticipantCreator: // Owner
				userStatus[user.UserID] = data.Owner
			case *tg.ChannelParticipantBanned: // Banned/kicked user
				userStatus[user.Peer.(*tg.PeerUser).UserID] = data.Banned
			case *tg.ChannelParticipantLeft: // A participant that left the channel/supergroup
				userStatus[user.Peer.(*tg.PeerUser).UserID] = data.Left
			default:
				t.log.Errorf("unexpected user type %T", user)
				return nil, errors.Errorf("unexpected user type %T", user)
			}
		}

		tgUsers = append(tgUsers, removeDuplicateUser(participantsObject.Users)...)

		for i := range participantsObject.Users {
			tgUser := participantsObject.Users[i].(*tg.User)
			if userStatus[tgUser.ID] == "" {
				continue
			}
			dbUsers = append(dbUsers, data.User{
				Username:    &tgUser.Username,
				Phone:       &tgUser.Phone,
				FirstName:   tgUser.FirstName,
				LastName:    tgUser.LastName,
				TelegramId:  tgUser.ID,
				AccessHash:  tgUser.AccessHash,
				AccessLevel: userStatus[tgUser.ID],
			})
		}

		//if we got all users with one api request, so we don't need to make more requests
		if len(participantsObject.Users) == participantsObject.Count {
			break
		}

		offset += limit
	}

	return dbUsers, nil
}

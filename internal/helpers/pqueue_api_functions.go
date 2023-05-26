package helpers

import (
	"strconv"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetChat(queue *pqueue.PriorityQueue, function any, args []any, priority int) (*tg_client.Chat, error) {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "some error while getting chat from api")
	}

	chat, ok := item.Response.Value.(*tg_client.Chat)
	if !ok {
		return nil, errors.Wrap(err, "wrong response type while getting chat from api")
	}

	return chat, nil
}

func GetChats(queue *pqueue.PriorityQueue, function any, args []any, priority int) ([]tg_client.Chat, error) {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "some error while getting chat from api")
	}

	chats, ok := item.Response.Value.([]tg_client.Chat)
	if !ok {
		return nil, errors.Wrap(err, "wrong response type while getting chat from api")
	}

	return chats, nil
}

func GetUser(queue *pqueue.PriorityQueue, function any, args []any, priority int) (*data.User, error) {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "some error while getting user from api")
	}

	user, ok := item.Response.Value.(*data.User)
	if !ok {
		return nil, errors.Wrap(err, "wrong response type while getting chat from api")
	}

	return user, nil
}

func GetUsers(queue *pqueue.PriorityQueue, function any, args []any, priority int) ([]data.User, error) {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "some error while getting users from api")
	}

	users, ok := item.Response.Value.([]data.User)
	if !ok {
		return nil, errors.Wrap(err, "wrong response type while getting chat users from api")
	}

	return users, nil
}

func GetRequestError(queue *pqueue.PriorityQueue, function any, args []any, priority int) error {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return err
	}

	return nil
}

func GetString(queue *pqueue.PriorityQueue, function any, args []any, priority int) (string, error) {
	item, err := AddFunctionInPQueue(queue, function, args, priority)
	if err != nil {
		return "", errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return "", err
	}

	myString, ok := item.Response.Value.(string)
	if !ok {
		return "", errors.Wrap(err, "wrong response type while getting chat users from api")
	}

	return myString, nil
}

func RetrieveChat(chats []tg_client.Chat, link string, id int64, accessHash *int64) *tg_client.Chat {
	if len(chats) == 1 {
		return &chats[0]
	}

	for i := range chats {
		if chats[i].Title != link {
			continue
		}

		if chats[i].Id != id {
			continue
		}

		if chats[i].AccessHash == nil && accessHash == nil {
			return &chats[i]
		}

		if chats[i].AccessHash != nil && accessHash != nil {
			if *(chats[i].AccessHash) == *(accessHash) {
				return &chats[i]
			}
		}
	}

	return nil
}

func ConvertIdentifiersStringsToInt(userIdStr, submoduleIdStr string, submoduleAccessHashStr *string) (userId, submoduleId int64, submoduleAccessHash *int64, err error) {
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse user id")
	}

	submoduleId, err = strconv.ParseInt(submoduleIdStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse submodule id")
	}

	if submoduleAccessHashStr == nil {
		return userId, submoduleId, nil, nil
	}

	tmp, err := strconv.ParseInt(*submoduleAccessHashStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse submodule id")
	}
	submoduleAccessHash = &tmp

	return userId, submoduleId, submoduleAccessHash, nil
}

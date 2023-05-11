package helpers

import (
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

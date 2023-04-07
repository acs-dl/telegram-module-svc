package processor

import (
	"container/heap"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) addFunctionInPqueue(function any, functionArgs []any, priority int) (*pqueue.QueueItem, error) {
	newUuid := uuid.New()
	queueItem := &pqueue.QueueItem{
		Uuid:     newUuid,
		Func:     function,
		Args:     functionArgs,
		Priority: priority,
	}
	heap.Push(p.pqueue, queueItem)
	item, err := p.pqueue.WaitUntilInvoked(newUuid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait until invoked")
	}

	err = p.pqueue.RemoveByUUID(newUuid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove by uuid")
	}

	return item, nil
}

func (p *processor) convertUserFromInterfaceAndCheck(userInterface any) (*data.User, error) {
	user, ok := userInterface.(*data.User)
	if !ok {
		return nil, errors.Errorf("wrong response type while getting users from api")
	}
	if user == nil {
		return nil, errors.Errorf("something wrong with user from api")
	}

	return user, nil
}

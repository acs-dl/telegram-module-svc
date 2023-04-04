package processor

import (
	"container/heap"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) addFunctionInPqueue(function any, functionArgs []any, priority int) *pqueue.QueueItem {
	newUuid := uuid.New()
	queueItem := &pqueue.QueueItem{
		Uuid:     newUuid,
		Func:     function,
		Args:     functionArgs,
		Priority: priority,
	}
	heap.Push(p.pqueue, queueItem)
	item := p.pqueue.WaitUntilInvoked(newUuid)
	p.pqueue.RemoveByUUID(newUuid)

	return item
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

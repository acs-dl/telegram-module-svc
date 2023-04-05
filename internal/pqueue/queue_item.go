package pqueue

import (
	"reflect"
	"sync"

	"github.com/google/uuid"
)

type ItemStatus string

const (
	PROCESSING ItemStatus = "in progress"
	INVOKED    ItemStatus = "invoked"
)

type Response struct {
	Value interface{}
	Error error
}
type QueueItem struct {
	Uuid uuid.UUID

	Func     interface{}
	Args     []interface{}
	Response Response

	Priority int
	index    int
	invoked  ItemStatus

	mu   sync.Mutex
	cond *sync.Cond
}

func (item *QueueItem) makeCall() {
	value := reflect.ValueOf(item.Func)

	args := make([]reflect.Value, 0)
	for _, arg := range item.Args {
		args = append(args, reflect.ValueOf(arg))
	}

	result := value.Call(args)

	if len(result) == 1 {
		item.Response.Error = checkIfError(result[0].Interface())
	} else {
		item.Response.Value = result[0].Interface()
		item.Response.Error = checkIfError(result[1].Interface())
	}
	item.invoked = INVOKED
}

func checkIfError(respErr interface{}) error {
	if err, ok := respErr.(error); ok {
		return err
	}
	return nil
}

func (item *QueueItem) waitInvoked() {
	item.mu.Lock()
	defer item.mu.Unlock()

	if item.cond == nil {
		item.cond = sync.NewCond(&item.mu)
	}

	for item.invoked != INVOKED {
		item.cond.Wait()
	}
}

func (item *QueueItem) callFunction() {
	item.mu.Lock()
	defer item.mu.Unlock()

	item.makeCall()

	if item.cond != nil {
		item.cond.Broadcast()
	}
}

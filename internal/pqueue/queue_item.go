package pqueue

import (
	"reflect"

	"github.com/google/uuid"
)

type Response struct {
	Value interface{}
	Error error
}
type QueueItem struct {
	Uuid     uuid.UUID
	Func     interface{}
	Args     []interface{}
	Response Response
	Priority int
	invoked  bool
	index    int
}

func (qi *QueueItem) callFunction() {
	value := reflect.ValueOf(qi.Func)

	args := make([]reflect.Value, 0)
	for _, arg := range qi.Args {
		args = append(args, reflect.ValueOf(arg))
	}

	result := value.Call(args)

	if len(result) == 1 {
		qi.Response.Error = checkIfError(result[0].Interface())
	} else {
		qi.Response.Value = result[0].Interface()
		qi.Response.Error = checkIfError(result[1].Interface())
	}
	qi.invoked = true
}

func checkIfError(respErr interface{}) error {
	if err, ok := respErr.(error); ok {
		return err
	}
	return nil
}

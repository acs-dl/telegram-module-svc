package helpers

import (
	"container/heap"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func AddFunctionInPQueue(pq *pqueue.PriorityQueue, function any, functionArgs []any, priority int) (*pqueue.QueueItem, error) {
	queueItem := &pqueue.QueueItem{
		Id:       GetFunctionSignature(function, functionArgs),
		Func:     function,
		Args:     functionArgs,
		Priority: priority,
	}
	heap.Push(pq, queueItem)

	item, err := pq.WaitUntilInvoked(queueItem.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait until invoked")
	}

	err = pq.RemoveById(queueItem.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove by id")
	}

	return item, nil
}

func GetFunctionName(function interface{}) string {
	splitName := strings.Split(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), ".")
	return splitName[len(splitName)-1]
}

func GetFunctionSignature(function interface{}, args []interface{}) string {
	signatureParts := []string{GetFunctionName(function), "("}

	signatureParts = append(signatureParts)

	for _, arg := range args {
		signatureParts = append(signatureParts, fmt.Sprintf("%v", arg))
	}

	signatureParts = append(signatureParts, ")")

	return strings.Join(signatureParts, " ")
}

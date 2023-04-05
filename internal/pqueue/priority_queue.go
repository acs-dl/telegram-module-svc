package pqueue

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type PriorityQueue []*QueueItem

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority > (*pq)[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	item.invoked = PROCESSING
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) RemoveByUUID(uuid uuid.UUID) error {
	item, err := pq.getElement(uuid)
	if err != nil {
		return err
	}
	*pq = append((*pq)[:item.index], (*pq)[item.index+1:]...)
	return nil
}

func (pq *PriorityQueue) getElement(uuid uuid.UUID) (*QueueItem, error) {
	for i := range *pq {
		if (*pq)[i].Uuid.String() == uuid.String() {
			return (*pq)[i], nil
		}
	}

	return nil, errors.New("element not found")
}

func (pq *PriorityQueue) WaitUntilInvoked(uuid uuid.UUID) (*QueueItem, error) {
	log.Printf("waiting until invoked for `%s`", uuid.String())

	item, err := pq.getElement(uuid)
	if err != nil {
		return nil, err
	}

	item.waitInvoked()

	return item, nil
}

func (pq *PriorityQueue) ProcessQueue(requestLimit int64, timeLimit time.Duration, stop chan struct{}) {
	ticker := time.NewTicker(timeLimit / time.Duration(requestLimit))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pq.processNextItem()
		case <-stop:
			return
		}
	}
}

func (pq *PriorityQueue) processNextItem() {
	for i := 0; i < pq.Len(); i++ {
		item := (*pq)[i]
		if item == nil {
			continue
		}

		if item.invoked == INVOKED {
			continue
		}

		item.callFunction()
		return
	}
}

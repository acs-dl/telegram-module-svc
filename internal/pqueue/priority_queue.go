package pqueue

import (
	"container/heap"
	"fmt"

	"github.com/google/uuid"
)

type QueueItem struct {
	Uuid     uuid.UUID
	Func     interface{}
	Priority int
	invoked  bool
	index    int
}

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

func (pq *PriorityQueue) getElement(uuid uuid.UUID) *QueueItem {
	for i := range *pq {
		if (*pq)[i].Uuid.String() == uuid.String() {
			return (*pq)[i]
		}
	}

	return nil
}

func (pq *PriorityQueue) isInvoked(uuid uuid.UUID) bool {
	element := pq.getElement(uuid)
	if element == nil {
		return false
	}

	return element.invoked
}

func TestHeap(items map[interface{}]int) {
	pq := make(PriorityQueue, 0)

	for value, priority := range items {
		item := &QueueItem{
			Func:     value,
			Priority: priority,
		}
		heap.Push(&pq, item)
	}
	//heap.Init(&pq)
	//
	//item := &QueueItem{
	//	Func:     "orange",
	//	priority: 1,
	//}
	//heap.Push(&pq, item)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*QueueItem)
		fmt.Printf("%.2d:%s ", item.Priority, item.Func)
	}
}

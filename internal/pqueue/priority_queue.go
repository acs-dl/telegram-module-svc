package pqueue

import (
	"fmt"
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
	item.invoked = false
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

func (pq *PriorityQueue) RemoveByUUID(uuid uuid.UUID) {
	item := pq.getElement(uuid)
	*pq = append((*pq)[:item.index], (*pq)[item.index+1:]...)
}

func (pq *PriorityQueue) getElement(uuid uuid.UUID) *QueueItem {
	for i := range *pq {
		if (*pq)[i].Uuid.String() == uuid.String() {
			return (*pq)[i]
		}
	}

	return nil
}

func (pq *PriorityQueue) WaitUntilInvoked(uuid uuid.UUID) *QueueItem {
	log.Printf("waiting until invoked for `%s`", uuid.String())
	for {
		element := pq.getElement(uuid)
		if element == nil {
			log.Printf("not found with uuid `%s`", uuid.String())
			return nil
		}

		if element.invoked {
			return element
		}

		log.Printf("need to sleep 3sec")
		time.Sleep(3 * time.Second)
	}
}

// ProcessQueue amount - the number of request for proper rate limit, duration - time to limit (e.g. 30, time.Second)
// TODO: ask about better way for limiting requests
func (pq *PriorityQueue) ProcessQueue(amount int64, duration time.Duration) {
	for {
		for i := 0; i < pq.Len(); i++ {
			item := (*pq)[i]
			if item == nil { //sometimes item can be nil that causes segfault
				continue
			}

			if item.invoked {
				continue
			}
			item.callFunction()

			fmt.Println(time.Duration(duration.Nanoseconds() / amount))
			time.Sleep(time.Duration(duration.Nanoseconds() / amount))
		}
	}
}

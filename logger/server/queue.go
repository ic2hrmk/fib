package server

import (
	"unsafe"
	"sync"
	"fmt"

	protocol "github.com/ic2hrmk/fib/common/contracts"
	"log"
)

type Queue struct {
	maxLength int64
	list      []protocol.Payload

	mutex sync.Mutex
}

func NewQueue(maxByteSize int64) *Queue {
	return &Queue{
		maxLength: maxByteSize / int64(unsafe.Sizeof(protocol.Payload{})),
	}
}

func (q Queue) Add(payload protocol.Payload) (err error) {
	if !q.HasFreeSpace() {
		err = fmt.Errorf("queue has no free space")
		return
	}

	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.list = append(q.list, payload)

	return
}

func (q Queue) Pop() (payload protocol.Payload, err error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.list) == 0 {
		err = fmt.Errorf("queue is empty")
		return
	}

	payload = q.list[len(q.list) - 1]
	q.list = q.list[:len(q.list) - 1]
	return
}

func (q Queue) HasFreeSpace() (bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return int64(len(q.list)) < q.maxLength
}
package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	data []interface{}
	lock sync.RWMutex
}

func (q *Queue) New() *Queue {
	q.data = []interface{}{}
	return q
}

// 此处是指针接收者
// 如果是值接收者，后面调用Enqueue()时，是对q的副本进行操作
func (q *Queue) Enqueue(item interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.Size() > 0 {
		q.Lock()
		defer q.Unlock()
		v := q.data[0]
		q.data = q.data[1:]
		return v, nil
	}
	return -1, errors.New("pop Error: Queue is empty")
}

func (q *Queue) Front() (interface{}, error) {
	if q.Size() > 0 {
		q.Lock()
		defer q.Unlock()
		return q.data[0], nil
	}
	return -1, errors.New("peep Error: Queue is empty")
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Lock() {
	q.lock.Lock()
}

func (q *Queue) Unlock() {
	q.lock.Unlock()
}

// ? 貌似切片可以存放同类型的数据

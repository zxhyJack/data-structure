package queue

import (
	"testing"
)

var q Queue

func initQueue() {
	if q.data == nil {
		q.New()
	}
}

func TestEnqueue(t *testing.T) {
	initQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if size := q.Size(); size != 3 {
		t.Errorf("wrong count, expectd 3 but got %d", size)
	}
}

func TestDequeue(t *testing.T) {
	q.Enqueue(1)
	q.Dequeue()
	if size := len(q.data); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

// func TestQueue(t *testing.T) {
// 	q := &Queue{
// 		data: []interface{}{1, 2, 3, "hello"},
// 	}

// 	t.Log(q.data)
// 	q.Enqueue(4)
// 	t.Log(q.data)

// 	qu := &Queue{
// 		data: make([]interface{}, 0),
// 	}

// 	t.Logf("Enque: A\n")
// 	qu.Enqueue("A")
// 	t.Logf("Enque B\n")
// 	qu.Enqueue("B")
// 	t.Logf("Len: %d\n", qu.Size())

// 	for qu.Size() > 0 {
// 		frontVal, _ := qu.Front()
// 		t.Logf("Front: %s\n", frontVal)
// 		v, err := qu.Dequeue()
// 		if err != nil {
// 			t.Log(err)
// 		}
// 		t.Logf("Dequeue: %s\n", v)
// 	}
// 	t.Logf("Len: %d\n", qu.Size())
// }

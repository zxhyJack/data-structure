package singleLinkedlist

import "testing"

func TestList(t *testing.T) {
	list := SingleLinkedList{
		head: &Node{data: 1},
		len:  1,
	}

	list.AddFront(2)
	list.AddFront(3)
	list.AddFront(4)
	list.AddFront(5)

	tmp := *list.head
	t.Log(tmp)

	t.Log(list.Traverse())
	// t.Log(list.Size())

	// if err := list.Insert(3, 1); err != nil {
	// 	t.Log(err)
	// }
	// t.Log(list.Traverse())

	// if _, err := list.RemoveAt(0); err != nil {
	// 	t.Log(err)
	// }
	// t.Log(list.Traverse())

	// t.Log(list.Indexof(5))

	// t.Log(list.Size2())
}

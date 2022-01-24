package doubleLinkedList

import "testing"

func TestDoubleLinkedList(t *testing.T) {
	d := DoubleLinkedList{}
	t.Log(d.Size())

	d.AddFront(1)
	d.AddFront(2)
	d.AddFront(3)

	t.Log(d.Traverse())

	d.AddBack(4)
	d.AddBack(5)

	t.Log(d.Traverse())
	t.Log(d.TraverseReverse())
	t.Log(d.Size())

	d.Insert(3, "a")
	d.Insert(0, "b")
	t.Log(d.Traverse())

	d.RemoveAt(0)
	t.Log(d.Traverse())

	d.RemoveAt(3)
	t.Log(d.Traverse())

	d.RemoveAt(4)
	t.Log(d.Traverse())
}

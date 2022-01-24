package doubleLinkedList

import "errors"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (d *DoubleLinkedList) AddFront(elem interface{}) {
	node := Node{data: elem}

	if d.head == nil {
		d.head = &node
		d.tail = &node
	} else {
		node.next = d.head
		d.head.prev = &node
		d.head = &node
	}

	d.len++
}

func (d *DoubleLinkedList) AddBack(elem interface{}) {
	node := Node{data: elem}

	if d.head == nil {
		d.head = &node
		d.tail = &node
	} else {
		node.prev = d.tail
		d.tail.next = &node
		d.tail = &node
	}

	d.len++
}

func (d *DoubleLinkedList) Traverse() []interface{} {
	list := []interface{}{}

	current := d.head
	for current != nil {
		list = append(list, current.data)
		current = current.next
	}

	return list
}

func (d *DoubleLinkedList) TraverseReverse() []interface{} {
	list := []interface{}{}

	current := d.tail
	for current != nil {
		list = append(list, current.data)
		current = current.prev
	}

	return list
}

func (d *DoubleLinkedList) Size() int {
	size := 0
	current := d.head
	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (d *DoubleLinkedList) Insert(i int, elem interface{}) error {
	if i < 0 || i > d.Size() {
		return errors.New("index out of bounds")
	}

	if i == 0 {
		d.AddFront(elem)
		return nil
	}

	prev := d.head
	for count := 0; count < i-1; count++ {
		prev = prev.next
	}

	node := Node{data: elem}
	current := prev.next

	node.next = current
	current.prev = &node

	node.prev = prev
	prev.next = &node

	d.len++
	return nil
}

func (d *DoubleLinkedList) RemoveAt(i int) error {
	if i < 0 || i > d.Size() {
		return errors.New("index out of bounds")
	}

	// 若移除的是链表的第一个元素
	if i == 0 {
		d.head = d.head.next
		d.head.prev = nil
		return nil
	}

	// 若移除的是链表的末尾元素
	if i == d.Size()-1 {
		d.tail = d.tail.prev
		d.tail.next = nil
		return nil
	}

	prev := d.head
	for count := 0; count < i-1; count++ {
		prev = prev.next
	}

	next := prev.next.next

	prev.next = next
	next.prev = prev

	d.len--
	return nil
}

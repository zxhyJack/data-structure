package singleLinkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type SingleLinkedList struct {
	head *Node
	len  int
}

func (s *SingleLinkedList) AddFront(elem interface{}) {
	node := &Node{data: elem}

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
}

func (s *SingleLinkedList) AddBack(elem interface{}) {
	node := &Node{data: elem}

	if s.head == nil {
		s.head = node
	} else {
		tmp := s.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = node
	}
	s.len++
}

func (s *SingleLinkedList) RemoveFront() error {
	if s.head == nil {
		return errors.New("list is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *SingleLinkedList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	var prev *Node
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
	s.len--
	return nil
}

func (s *SingleLinkedList) Front() (interface{}, error) {
	if s.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	return s.head.data, nil
}

func (s *SingleLinkedList) Size() int {
	return s.len
}

func (s *SingleLinkedList) Traverse() []interface{} {
	if s.head == nil {
		return nil
	}

	items := make([]interface{}, 0)
	current := s.head
	// for current.next != nil{
	// 	fmt.Println(current.data)
	// 	current = current.next
	// }

	for current != nil {
		items = append(items, current.data)
		current = current.next
	}

	return items
}

func (s *SingleLinkedList) Insert(i int, elem interface{}) error {
	if i < 0 || i > s.Size() {
		return errors.New("index out of bounds")
	}
	node := Node{data: elem}

	if i == 0 {
		node.next = s.head
		s.head = &node
	} else {
		// count := 0
		// var prev *Node
		// current := s.head
		// for count != i {
		// 	prev = current
		// 	current = current.next
		// 	count++
		// }
		// node.next = current
		// prev.next = &node

		prev := s.head
		for count := 0; count < i-1; count++ {
			prev = prev.next
		}
		node.next = prev.next
		prev.next = &node
		s.len++
	}
	return nil
}

func (s *SingleLinkedList) RemoveAt(i int) (interface{}, error) {
	if i < 0 || i > s.len {
		return nil, fmt.Errorf("index out of bounds")
	}

	var prev *Node
	current := s.head
	if i == 0 {
		s.head = s.head.next
	} else {
		for count := 0; count < i; count++ {
			prev = current
			current = current.next

		}
		// current.next = current.next.next
		prev.next = current.next
	}

	s.len--
	return current, nil
}

func (s *SingleLinkedList) Indexof(elem interface{}) int {
	current := s.head
	for i := 0; i < s.len; i++ {
		if current.data == elem {
			return i
		}
		current = current.next
	}

	// j := 0
	// for {
	//     if node.content == t {
	//         return j
	//     }
	//     if node.next == nil {
	//         return -1
	//     }
	//     node = node.next
	//     j++
	// }

	return -1
}

func (s *SingleLinkedList) IsEmpty() bool {
	return s.head == nil
}

func (s *SingleLinkedList) Size2() int {
	count := 0
	current := s.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

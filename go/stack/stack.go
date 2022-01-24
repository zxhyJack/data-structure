package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	data []interface{}
	lock sync.RWMutex
}

func (s *Stack) New() *Stack {
	if s.data == nil {
		s.data = []interface{}{}
	}
	return s
}

func (s *Stack) Push(elem interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Size() > 0 {
		s.lock.Lock()
		defer s.lock.Unlock()

		v := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1] // 此处时len()-1，包含下限，不包含上限
		return v, nil
	}
	return -1, errors.New("pop error: stack is empty")
}

func (s *Stack) Front() (interface{}, error) {
	if s.Size() > 0 {
		s.lock.RLock()
		defer s.lock.RUnlock()
		return s.data[s.Size()-1], nil
	}
	return -1, errors.New("peer error: stack is empty")
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

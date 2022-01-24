package set

import (
	"errors"
	"sync"
)

type Set struct {
	data map[interface{}]struct{}
	lock sync.RWMutex
}

func (s *Set) Add(elem interface{}) *Set {
	s.lock.Lock()
	defer s.lock.Unlock()

	// 如果set只声明未初始化，data需要初始化
	if s.data == nil {
		s.data = make(map[interface{}]struct{})
	}

	// 如果元素已经存在直接返回，未存在则添加元素
	if !s.Exists(elem) {
		s.data[elem] = struct{}{}
	}

	return s
}

func (s *Set) Remove(elem interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Exists(elem) {
		delete(s.data, elem)
		return nil
	} else {
		return errors.New("remove error: item doesn't exist in the set")
	}
}

func (s *Set) Exists(elem interface{}) bool {
	_, ok := s.data[elem]
	return ok
}

func (s *Set) Items() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	r := []interface{}{}
	for k := range s.data {
		r = append(r, k)
	}
	return r
}
func (s *Set) Clear() *Set {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = make(map[interface{}]struct{})
	return s
}

func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.data)
}

func (s *Set) Union(s2 *Set) *Set {
	s3 := &Set{}

	s.lock.RLock()
	for _, item := range s.Items() {
		s3.Add(item)
	}
	s.lock.RUnlock()

	s2.lock.RLock()
	for _, item := range s2.Items() {
		if ok := s3.Exists(item); !ok {
			s3.Add(item)
		}
	}
	s2.lock.RUnlock()

	return s3
}

func (s *Set) Intersection(s2 *Set) *Set {
	s3 := Set{}

	s.lock.RLock()
	s2.lock.RLock()
	for _, item := range s.Items() {
		if ok := s2.Exists(item); ok {
			s3.Add(item)
		}

		// if _, ok := s2.data[item]; ok {
		// 	s3.data[item] = struct{}{}
		// }

		// for _, y := range s2.Items() {
		// 	if x == y {
		// 		s3.Add(x)
		// 	}
		// }
	}
	s.lock.RUnlock()
	s2.lock.RUnlock()

	return &s3
}

func (s *Set) Difference(s2 *Set) *Set {
	// s3 := Set{} // data仍是零值nil
	s3 := Set{
		data: make(map[interface{}]struct{}),
	}

	// s.lock.RLock()
	// for _, item := range s.Items() {
	// 	s3.Add(item)
	// }
	// s.lock.RUnlock()

	// s2.lock.RLock()
	// for _, item := range s2.Items() {
	// 	if ok := s3.Exists(item); ok {
	// 		s3.Remove(item)
	// 	}
	// }
	// s2.lock.RUnlock()

	for _, item := range s.Items() {
		// if ok := s2.Exists(item); !ok {
		// 	s3.Add(item)
		// }

		if _, ok := s2.data[item]; !ok {
			s3.data[item] = struct{}{}
		}
	}

	return &s3
}

func (s *Set) Subset(s2 *Set) bool {
	s.lock.RLock()
	s2.lock.RLock()
	defer s.lock.RUnlock()
	defer s2.lock.RUnlock()

	for _, item := range s.Items() {
		if ok := s2.Exists(item); !ok {
			return false
		}
	}

	return true
}

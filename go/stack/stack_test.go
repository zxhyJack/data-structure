package stack

import "testing"

var s Stack

func TestPush(t *testing.T) {
	s = *s.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	if size := len(s.data); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestPop(t *testing.T) {
	s.Pop()
	if v, err := s.Pop(); err != nil {
		t.Errorf(err.Error())
	} else {
		if v != 2 {
			t.Errorf("wrong pop, expected 2 but got %d", v)
		}
	}
	s.Pop()
	if size := len(s.data); size != 0 {
		t.Errorf("wrong count, expected 0 but got %d", size)
	}
}

func TestFront(t *testing.T) {
	if v, err := s.Front(); err != nil {
		t.Log(err.Error())
	} else {
		t.Log(v)
	}
}

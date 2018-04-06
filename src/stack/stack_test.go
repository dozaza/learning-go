package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := new(Stack)
	s.push(1)
	if s.size != 1 {
		t.Log(s.size)
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := new(Stack)
	s.push(10)
	s.push(11)
	poped := s.pop()
	if poped != 11 {
		t.Logf("Wrong poped: %d" , poped)
		t.Fail()
	}
}

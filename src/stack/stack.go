package stack

import "strconv"

// Q15

type Stack struct {
	size int
	data [5]int
}

func (s Stack) String() string {
	if s.size == 0 {
		return ""
	}

	str := ""
	for i, d := range s.data {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(d) + "] "
	}
	return str
}

func (s *Stack) push(num int) {
	if s.size < 5 {
		s.data[s.size] = num
		s.size += 1
	}
}

func (s *Stack) pop() int {
	if s.size == 0 {
		panic("empty stack")
	}

	s.size -= 1
	return s.data[s.size]
}
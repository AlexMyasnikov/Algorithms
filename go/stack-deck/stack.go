package stackdeck

import (
	"errors"
	"fmt"
)

type Stack struct {
	s []float64
}

var errEmpty = errors.New("stack is empty")

func (s *Stack) Push(el float64) {
	s.s = append(s.s, el)
}

func (s *Stack) Pop() (float64, error) {
	if len(s.s) == 0 {
		return 0, errEmpty
	}
	last := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return last, nil
}

func (s *Stack) Pick() (float64, error) {
	if len(s.s) == 0 {
		return 0, errEmpty
	}
	return s.s[len(s.s)-1], nil
}

func (s *Stack) Show() {
	for i := len(s.s) - 1; i >= 0; i-- {
		fmt.Println(s.s[i])
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

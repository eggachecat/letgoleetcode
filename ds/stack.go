package ds

import "errors"

// Stack blabla
type Stack struct {
	s []string
}

// Push blabla
func (s *Stack) Push(v string) {
	s.s = append(s.s, v)
}

// Pop blabla
func (s *Stack) Pop() (string, error) {
	l := len(s.s)
	if l == 0 {
		return "ERROR", errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

// IsEmpty blabla
func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

// NewStack blabla
func NewStack() *Stack {
	return &Stack{make([]string, 0)}
}

package ds

import "errors"

// Queue blabla
type Queue struct {
	s []string
}

// Push blabla
func (s *Queue) Push(v string) {
	s.s = append(s.s, v)
}

// Pop blabla
func (s *Queue) Pop() (string, error) {
	l := len(s.s)
	if l == 0 {
		return "ERROR", errors.New("Empty Stack")
	}

	res := s.s[0]
	s.s = s.s[1:]
	return res, nil
}

// IsEmpty blabla
func (s *Queue) IsEmpty() bool {
	return len(s.s) == 0
}

// NewQueue blabla
func NewQueue() *Queue {
	return &Queue{make([]string, 0)}
}

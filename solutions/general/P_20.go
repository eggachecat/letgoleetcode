package general

import (
	"lc/ds"
)

// Problem20 is for valid-parentheses
func Problem20(s string) bool {
	m := map[string]int{
		"{": 1, "[": 2, "(": 3,
		"}": -1, "]": -2, ")": -3,
	}

	stack := ds.NewStack()
	for _, char := range s {
		str := string(char)
		o := m[str]
		if o < 0 {
			pairStr, err := stack.Pop()
			if err != nil {
				return false
			}
			if m[pairStr]+o != 0 {
				return false
			}
		} else {
			stack.Push(str)
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}

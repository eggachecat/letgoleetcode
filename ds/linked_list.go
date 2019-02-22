package ds

type LinkedList struct {
	Next  *LinkedList
	Value interface{}
}

func (l *LinkedList) Insert(value interface{}) {
	ptr := l
	for ptr != nil {
		ptr = ptr.Next
	}

	ptr.Next = &LinkedList{
		Value: value,
	}
}

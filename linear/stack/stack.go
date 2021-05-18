package stack

import (
	"foundation/linear/linked_list/double"
)

// 栈
type Stack struct {
	linkedList *double.LinkedList
}

func (s *Stack) Size() int {
	return s.linkedList.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(element interface{}) error {
	return s.linkedList.Add(element)
}

func (s *Stack) Pop() (interface{}, error) {
	 return s.linkedList.RemoveLast()
}

func (s *Stack) Peek() (interface{}, bool) {
	return s.linkedList.Get(s.linkedList.Size() - 1)
}

func (s *Stack) Clear() {
	s.linkedList.RemoveAll()
}
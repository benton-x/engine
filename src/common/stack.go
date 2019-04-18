package common

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

// NewStack ...
func NewStack() *Stack {
	l := list.New()
	return &Stack{l}
}

// Push add elem to stack
func (s *Stack) Push(elem interface{}) {
	s.list.PushBack(elem)
}

// Pop return the top elem from stack
func (s *Stack) Pop() interface{} {
	elem := s.list.Back()
	if elem == nil {
		return nil
	}
	s.list.Remove(elem)
	return elem.Value
}

// Peak return the value of the top elem
func (s *Stack) Peak() interface{} {
	elem := s.list.Back()
	if elem == nil {
		return nil
	}
	return elem.Value
}

// Len return the length of the stack
func (s *Stack) Len() int {
	return s.list.Len()
}

// Empty check the length of the stack
func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

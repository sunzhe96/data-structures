package stack

import (
	"github.com/sunzhe96/data-structures/list"
)

type Stack struct {
	top list.List
}

func New() *Stack {
	s := Stack{}
	return &s
}

func (s *Stack) Push(key int) {
	s.top.PushFront(key)
}

func (s *Stack) Pop() int {
	return s.top.PopFront()
}

func (s *Stack) Top() int {
	return s.top.TopFront()
}

func (s *Stack) Empty() bool {
	return s.top.Empty()
}

func (s *Stack) Print() {
	s.top.Print()
}

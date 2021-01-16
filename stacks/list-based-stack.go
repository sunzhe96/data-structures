package stacks

import (
	"github.com/sunzhe96/data-structures/linked-lists"
)

// This implementation is based on my singly-linked lists library.

// | API     | function                                    |
// |         |                                             |
// | Push()  | adds key to collection                      |
// |         |                                             |
// | Pop()   | removes and returns most recently-added key |
// |         |                                             |
// | Top()   | returns most recently-added key             |
// |         |                                             |
// | Empty() | are there any elements?                     |

/*
   list implementation

type node struct {
	key  int
	next *node
}

type List struct {
	head *node
}
*/

type ListStack struct {
	top lists.List
}

func (s *ListStack) Push(key int) {
	s.top.PushFront(key)
}

func (s *ListStack) Pop() int {
	return s.top.PopFront()
}

func (s *ListStack) Top() int {
	return s.top.TopFront()
}

func (s *ListStack) Empty() bool {
	return s.top.Empty()
}

package stack

import "fmt"

type SliceStack []int

func NewSliceStack() SliceStack {
	s := SliceStack{}
	return s
}

func (s *SliceStack) Push(key int) {
	*s = append(*s, key)
}

func (s *SliceStack) Pop() int {
	arr := *s
	res := arr[len(arr)-1]
	*s = arr[:len(arr)-1]
	return res
}

func (s *SliceStack) Top() int {
	arr := *s
	if arr.Empty() {
		panic("empty stack")
	}
	return arr[len(arr)-1]
}

func (s *SliceStack) Empty() bool {
	arr := *s
	if len(arr) == 0 {
		return true
	}
	return false
}

func (s *SliceStack) Print() {
	arr := *s
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

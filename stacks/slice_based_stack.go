package stacks

// This is an implementation of stacks based on go's built-in data structure slice

// | API     | function                                    |
// |         |                                             |
// | push()  | adds key to collection                      |
// |         |                                             |
// | pop()   | removes and returns most recently-added key |
// |         |                                             |
// | top()   | returns most recently-added key             |
// |         |                                             |
// | empty() | are there any elements?                     |

type Stack []int

func (s *Stack) Push(key int) {
	*s = append(*s, key)
}

func (s *Stack) Pop() int {
	arr := *s
	res := arr[len(arr)-1]
	*s = arr[:len(arr)-1]
	return res
}

func (s Stack) Top() int {
	if s.Empty() {
		panic("empty stack")
	}
	return s[len(s)-1]
}

func (s Stack) Empty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

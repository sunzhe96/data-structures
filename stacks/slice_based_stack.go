package stack

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

type stack []int

func (s *stack) push(key int) {
	*s = append(*s, key)
}

func (s *stack) pop() int {
	arr := *s
	res := arr[len(arr)-1]
	*s = arr[:len(arr)-1]
	return res
}

func (s stack) top() int {
	if s.empty() {
		panic("empty stack")
	}
	return s[len(s)-1]
}

func (s stack) empty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

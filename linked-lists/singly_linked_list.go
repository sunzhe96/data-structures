package lists

import "fmt"

// | API                  | function                     |
// |----------------------|------------------------------|
// | PushFront(Key)       | add key to front             |
// |                      |                              |
// | TopFront()           | return key of the front item |
// |                      |                              |
// | PopFront()           | remove the front item        |
// |                      |                              |
// | PushBack()           | add item to back             |
// |                      |                              |
// | TopBack()            | return key of back item      |
// |                      |                              |
// | Popback()            | remove back item             |
// |                      |                              |
// | Find(Key)            | is key in list?              |
// |                      |                              |
// | Erase(Key)           | remove key from list         |
// |                      |                              |
// | Empty()              | is list empty?               |
// |                      |                              |
// | PrintList()          | print the list               |

type node struct {
	key  int
	next *node
}

type List struct {
	head *node
}

func (l *List) PushFront(key int) {
	ptr := l.head
	newNode := node{
		key,
		ptr,
	}
	l.head = &newNode
	fmt.Printf("pushed %v to the front\n", key)
}

func (l *List) PushBack(key int) {
	if l.Empty() {
		fmt.Println("empty list, use pushFront instead")
		l.PushFront(key)
		return
	}

	var newNode node
	newNode.key = key

	ptr := l.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &newNode
}

func (l *List) PopFront() int {
	if l.Empty() {
		panic("error: empty list")

	}
	newHead := l.head.next
	res := l.head.key
	l.head = newHead
	return res
}

func (l *List) PopBack() int {
	if l.Empty() {
		panic("error: empty list")
	}
	ptr := l.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	res := ptr.next.key
	ptr.next = nil
	return res
}

func (l List) TopFront() int {
	if l.Empty() {
		panic("error: empty list")
	}
	fmt.Printf("top front is: %v\n", l.head.key)
	return l.head.key
}

func (l List) TopBack() int {
	if l.Empty() {
		panic("error: empty list")

	}
	var ptr, previous *node
	ptr = l.head

	for ptr != nil {
		previous = ptr
		ptr = ptr.next
	}
	fmt.Printf("top back is: %v\n", previous.key)
	return previous.key
}

func (l List) Find(key int) bool {
	ptr := l.head
	for ptr != nil {
		if ptr.key == key {
			return true
		}
		ptr = ptr.next
	}
	return false
}

func (l *List) Erase(key int) {
	ptr := l.head
	for ptr.next.next != nil {
		if ptr.next.key == key {
			ptr.next = ptr.next.next
		}
		ptr = ptr.next
	}
}

func (l List) Empty() bool {
	return l.head == nil
}

func (l List) PrintList() {
	if l.Empty() {
		fmt.Println("error: empty list")
		return
	}
	ptr := l.head
	for ptr != nil {
		fmt.Printf("%v ", ptr.key)
		ptr = ptr.next
	}
	fmt.Printf("\n")
}

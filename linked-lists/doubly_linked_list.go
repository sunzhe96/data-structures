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

type doublyNode struct {
	key  int
	next *doublyNode
	prev *doublyNode
}

type DoublyList struct {
	head *doublyNode
	tail *doublyNode
}

func (l *DoublyList) PushFront(key int) {
	newNode := doublyNode{
		key:  key,
		next: l.head,
		prev: nil,
	}

	// push to an empty list
	if l.Empty() {
		l.head = &newNode
		l.tail = &newNode
		fmt.Printf("pushed %v to the back\n", key)
		return
	}
	// push to a non-empty list
	l.head.prev = &newNode
	l.head = &newNode

	fmt.Printf("pushed %v to the front\n", key)
}

func (l *DoublyList) PushBack(key int) {
	newNode := doublyNode{
		key:  key,
		next: nil,
		prev: l.tail,
	}

	// push back to an empty list
	if l.Empty() {
		fmt.Println("empty list, use pushFront instead")
		l.PushFront(key)
		return
	}

	// push back to an non-empty list
	l.tail.next = &newNode
	l.tail = &newNode
	fmt.Printf("pushed %v to the back\n", key)
}

func (l *DoublyList) PopFront() int {
	if l.Empty() {
		panic("error: empty list")
	}
	newHead := l.head.next
	res := l.head.key
	l.head = newHead
	return res
}

func (l *DoublyList) PopBack() int {
	if l.Empty() {
		panic("error: empty list")
	}
	res := l.tail.prev.next.key
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	return res
}

func (l DoublyList) TopFront() int {
	if l.Empty() {
		panic("error: empty list")
	}
	fmt.Printf("top front is: %v\n", l.head.key)
	return l.head.key
}

func (l DoublyList) TopBack() int {
	if l.Empty() {
		panic("error: empty list")
	}
	fmt.Printf("top back is: %v\n", l.tail.key)
	return l.tail.key
}

func (l DoublyList) Find(key int) bool {
	ptr := l.head
	for ptr != nil {
		if ptr.key == key {
			return true
		}
		ptr = ptr.next
	}
	return false
}

func (l *DoublyList) Erase(key int) {
	ptr := l.head
	for ptr.next.next != nil {
		if ptr.next.key == key {
			ptr.next = ptr.next.next
		}
		ptr = ptr.next
	}
}

func (l DoublyList) Empty() bool {
	return l.head == nil
}

func (l DoublyList) PrintList() {
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

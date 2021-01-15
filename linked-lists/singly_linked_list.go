package singlylists

import "fmt"

// | API                  | function                     |
// |----------------------+------------------------------|
// | pushFront(Key)       | add key to front             |
// |                      |                              |
// | topFront()           | return key of the front item |
// |                      |                              |
// | popFront()           | remove the front item        |
// |                      |                              |
// | pushBack()           | add item to back             |
// |                      |                              |
// | topBack()            | return key of back item      |
// |                      |                              |
// | popback()            | remove back item             |
// |                      |                              |
// | find(Key)            | is key in list?              |
// |                      |                              |
// | erase(Key)           | remove key from list         |
// |                      |                              |
// | empty()              | is list empty?               |
// |                      |                              |
// | printList()          | print the list               |

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

func (l *List) PopFront() {
	if l.Empty() {
		fmt.Println("error: empty list")
		return
	}
	newHead := l.head.next
	l.head = newHead
}

func (l *List) PopBack() {
	if l.Empty() {
		fmt.Println("error: empty list")
		return
	}
	ptr := l.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	ptr.next = nil
}

func (l List) topFront() {
	if l.Empty() {
		fmt.Println("error: empty list")
		return
	}
	fmt.Printf("top front is: %v\n", l.head.key)
}

func (l List) Topback() {
	if l.Empty() {
		fmt.Println("error: empty list")
		return
	}
	var ptr, previous *node
	ptr = l.head

	for ptr != nil {
		previous = ptr
		ptr = ptr.next
	}
	fmt.Printf("top back is: %v\n", previous.key)
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
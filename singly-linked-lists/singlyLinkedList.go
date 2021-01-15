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

type list struct {
	head *node
}

func (l *list) pushFront(key int) {
	ptr := l.head
	newNode := node{
		key,
		ptr,
	}
	l.head = &newNode
	fmt.Printf("pushed %v to the front\n", key)
}

func (l *list) pushBack(key int) {
	if l.empty() {
		fmt.Println("empty list, use pushFront instead")
		l.pushFront(key)
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

func (l *list) popFront() {
	if l.empty() {
		fmt.Println("error: empty list")
		return
	}
	newHead := l.head.next
	l.head = newHead
}

func (l *list) popBack() {
	if l.empty() {
		fmt.Println("error: empty list")
		return
	}
	ptr := l.head
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	ptr.next = nil
}

func (l list) topFront() {
	if l.empty() {
		fmt.Println("error: empty list")
		return
	}
	fmt.Printf("top front is: %v\n", l.head.key)
}

func (l list) topBack() {
	if l.empty() {
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

func (l list) find(key int) bool {
	ptr := l.head
	for ptr != nil {
		if ptr.key == key {
			return true
		}
		ptr = ptr.next
	}
	return false
}

func (l *list) erase(key int) {
	ptr := l.head
	for ptr.next.next != nil {
		if ptr.next.key == key {
			ptr.next = ptr.next.next
		}
		ptr = ptr.next
	}
}

func (l list) empty() bool {
	return l.head == nil
}

func (l list) printList() {
	if l.empty() {
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

package doublylists

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
	prev *node
}

type list struct {
	head *node
	tail *node
}

func (l *list) pushFront(key int) {
	newNode := node{
		key:  key,
		next: l.head,
		prev: nil,
	}

	// push to an empty list
	if l.empty() {
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

func (l *list) pushBack(key int) {
	newNode := node{
		key:  key,
		next: nil,
		prev: l.tail,
	}

	// push back to an empty list
	if l.empty() {
		fmt.Println("empty list, use pushFront instead")
		l.pushFront(key)
		return
	}

	// push back to an non-empty list
	l.tail.next = &newNode
	l.tail = &newNode
	fmt.Printf("pushed %v to the back\n", key)
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

	l.tail.prev.next = nil
	l.tail = l.tail.prev
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

	fmt.Printf("top back is: %v\n", l.tail.key)
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

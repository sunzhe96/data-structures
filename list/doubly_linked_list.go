package list

import "fmt"

type doublyNode struct {
	key  int
	next *doublyNode
	prev *doublyNode
}

type DoublyList struct {
	head *doublyNode
	tail *doublyNode
}

func NewDoublyList() *DoublyList {
	l := DoublyList{}
	return &l
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
		return
	}
	// push to a non-empty list
	l.head.prev = &newNode
	l.head = &newNode
}

func (l *DoublyList) PushBack(key int) {
	newNode := doublyNode{
		key:  key,
		next: nil,
		prev: l.tail,
	}

	// push back to an empty list
	if l.Empty() {
		l.PushFront(key)
		return
	}

	// push back to an non-empty list
	l.tail.next = &newNode
	l.tail = &newNode
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

	return l.head.key
}

func (l DoublyList) TopBack() int {
	if l.Empty() {
		panic("error: empty list")
	}

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

func (l DoublyList) Print() {
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

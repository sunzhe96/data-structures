// package implement hash tables
package hashtable

import (
	"fmt"
)

// | API                | Operation                                        |
// |--------------------+--------------------------------------------------|
// | Search(Key)        | return the corresponding value                   |
// |                    |                                                  |
// | Insert(Key, Value) | insert a key-value pair in the dictionary        |
// |                    |                                                  |
// | Delete(Key)        | delete the key-value pair                        |
// |                    |                                                  |
// | MakeHashTable()    | initialize a hashtable and return the hash table |
// |                    |                                                  |
// | PrintTable()       | print the hash table by each slot                |

const arraySize = 10

type HashTable struct {
	array [arraySize]*slot
}

type slot struct {
	head *node
}

type node struct {
	key   string
	value int
	next  *node
}

func MakeHashTable() HashTable {
	table := &HashTable{}
	for i := range (*table).array {
		(*table).array[i] = &slot{}
	}
	return *table
}

func (h *HashTable) PrintTable() {
	for i := range h.array {
		ptr := h.array[i].head
		fmt.Printf("slot %d: ", i)
		for ptr != nil {
			fmt.Printf("%s : %d | ", ptr.key, ptr.value)
			ptr = ptr.next
		}
		fmt.Printf("\n")
	}
}

func (h *HashTable) Insert(key string, value int) {
	hashCode := hash(key)
	h.array[hashCode].pushFront(key, value)
}

func (h HashTable) Search(key string) {
	hashCode := hash(key)
	value, found := h.array[hashCode].find(key)
	if found == false {
		fmt.Println(key, " not found")
		return
	}
	fmt.Println(key, ": ", value)
}

func (h *HashTable) Delete(key string) {
	hashCode := hash(key)
	h.array[hashCode].delete(key)
}

func (s *slot) pushFront(key string, value int) {
	ptr := s.head
	newNode := &node{
		key,
		value,
		ptr,
	}
	s.head = newNode
}

func (s *slot) delete(key string) {
	ptr := s.head
	if ptr.key == key {
		s.head = ptr.next
		return
	}

	for ptr.next != nil {
		if ptr.next.key == key {
			ptr.next = ptr.next.next
		}
		ptr = ptr.next
	}
}

func (s *slot) find(key string) (int, bool) {
	ptr := s.head
	for ptr != nil {
		if ptr.key == key {
			return ptr.value, true
		}
		ptr = ptr.next
	}
	return 0, false
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

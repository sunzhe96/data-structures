package quene

import (
	"github.com/sunzhe96/data-structures/list"
)

type Quene struct {
	list list.DoublyList
}

func New() *Quene {
	q := Quene{}
	return &q
}

func (q *Quene) EnQuene(key int) {
	q.list.PushBack(key)
}

func (q *Quene) DeQuene() int {
	return q.list.PopFront()
}

func (q *Quene) Empty() bool {
	return q.list.Empty()
}

func (q *Quene) Print() {
	q.list.Print()
}

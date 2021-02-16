package quene

import (
	"github.com/sunzhe96/data-structures/linked-lists"
)

// | API       | function                                       |
// |-----------|------------------------------------------------|
// | EnQuene() | Adds key to the collection                     |
// |           |                                                |
// | DeQuene() | removes and returns the least recent-added key |
// |           |                                                |
// | Empty()   | are there any elements?                        |

type Quene struct {
	list lists.DoublyList
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

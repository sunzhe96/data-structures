package tree

import "fmt"

type BST struct {
	Root *Node
}

type Node struct {
	Key    int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(k int) *Node {
	n := Node{}
	n.Key = k
	return &n
}

func NewBST() *BST {
	t := BST{}
	return &t
}

func (t *BST) InorderTreeWalk() {
	ptr := t.Root
	if ptr != nil {
		leftSub := BST{ptr.Left}
		rightSub := BST{ptr.Right}
		leftSub.InorderTreeWalk()
		fmt.Printf("%d ", ptr.Key)
		rightSub.InorderTreeWalk()
	}
}

func (t *BST) PreorderTreeWalk() {
	ptr := t.Root
	if ptr != nil {
		leftSub := BST{ptr.Left}
		rightSub := BST{ptr.Right}
		fmt.Printf("%d ", ptr.Key)
		leftSub.PreorderTreeWalk()
		rightSub.PreorderTreeWalk()
	}
}

func (t *BST) PostorderTreeWalk() {
	ptr := t.Root
	if ptr != nil {
		leftSub := BST{ptr.Left}
		rightSub := BST{ptr.Right}
		leftSub.PostorderTreeWalk()
		rightSub.PostorderTreeWalk()
		fmt.Printf("%d ", ptr.Key)
	}
}

func (t *BST) Print() {
	t.InorderTreeWalk()
	fmt.Printf("\n")
}

func (t *BST) Search(k int) *Node {
	ptr := t.Root
	if ptr == nil || ptr.Key == k {
		return ptr
	}
	leftSub := BST{ptr.Left}
	rightSub := BST{ptr.Right}

	if k < ptr.Key {
		return leftSub.Search(k)
	}

	return rightSub.Search(k)
}

func (t *BST) SearchIter(k int) *Node {
	ptr := t.Root
LOOP:
	for ptr != nil {
		switch {
		case k < ptr.Key:
			ptr = ptr.Left
		case k > ptr.Key:
			ptr = ptr.Right
		default:
			break LOOP
		}
	}
	return ptr
}

func (t *BST) Min() *Node {
	ptr := t.Root
	if ptr.Left == nil {
		return ptr
	}
	leftSub := BST{ptr.Left}
	return leftSub.Min()
}

func (t *BST) MinIter() *Node {
	ptr := t.Root
	for ptr.Left != nil {
		ptr = ptr.Left
	}
	return ptr
}

func (t *BST) Max() *Node {
	ptr := t.Root
	if ptr.Right == nil {
		return ptr
	}
	rightSub := BST{ptr.Right}
	return rightSub.Max()
}

func (t *BST) MaxIter() *Node {
	ptr := t.Root
	for ptr.Right != nil {
		ptr = ptr.Right
	}
	return ptr
}

func (n *Node) Successor() *Node {
	ptr := n
	if ptr.Right != nil {
		sub := BST{ptr.Right}
		return sub.Min()
	}

	for ptr.Parent != nil && ptr.Parent.Right == ptr {
		ptr = ptr.Parent
	}

	return ptr.Parent
}

func (t *Node) Predecessor() *Node {
	ptr := t
	if ptr.Left != nil {
		sub := BST{ptr.Left}
		return sub.Max()
	}

	for ptr.Parent != nil && ptr.Parent.Left == ptr {
		ptr = ptr.Parent
	}

	return ptr.Parent
}

func (t *BST) Insert(n *Node) {
	if t.Root == nil {
		t.Root = n
		return
	}
	ptr := t.Root
	for ptr != nil {
		if n.Key < ptr.Key {
			n.Parent = ptr
			ptr = ptr.Left
		} else {
			n.Parent = ptr
			ptr = ptr.Right
		}
	}
	if n.Key < n.Parent.Key {
		n.Parent.Left = n
	} else {
		n.Parent.Right = n
	}
}

func (t *BST) transplant(u *Node, v *Node) {
	switch {
	case u.Parent == nil:
		t.Root = v
	case u == u.Parent.Left:
		u.Parent.Left = v
	case u == u.Parent.Right:
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func (t *BST) Delete(n *Node) {
	if n.Left == nil {
		t.transplant(n, n.Right)
	} else if n.Right == nil {
		t.transplant(n, n.Left)
	} else {
		rightSub := BST{n.Right}
		newNode := rightSub.Min()
		if newNode.Parent != n {
			t.transplant(newNode.Right, n.Right)
		}
		t.transplant(n, newNode)
		newNode.Left = n.Left
		newNode.Left.Parent = newNode
	}
}

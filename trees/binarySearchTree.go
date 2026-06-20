package trees

// Binary search tree

import (
	"fmt"
)

type NumericTypes interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Node[T NumericTypes] struct {
	Left  *Node[T]
	Right *Node[T]
	Data  T
}

// BSTree() -> Binary search tree, &BSTree[type] , must be in NumericTypes ;
type BSTree[T NumericTypes] struct {
	Root *Node[T]
}

// Insert() -> Inserts Node
func (tree *BSTree[T]) Insert(n T) {
	tree.Root = insertHelper(tree.Root, n)
}

func insertHelper[T NumericTypes](node *Node[T], n T) *Node[T] {
	if node == nil {
		return &Node[T]{Data: n}
	}

	switch {
	case n < node.Data:
		node.Left = insertHelper(node.Left, n)

	case n > node.Data:
		node.Right = insertHelper(node.Right, n)
	}

	return node
}

// Display() -> prints nodes :: recursive implementation
func (t *BSTree[T]) Display() {
	displayHelper(t.Root)
}

func displayHelper[T NumericTypes](node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Println("node:", node.Data)
	if node.Left != nil {
		fmt.Println("  left child:", node.Left.Data)
	}
	if node.Right != nil {
		fmt.Println("  right child:", node.Right.Data)
	}

	displayHelper(node.Left)
	displayHelper(node.Right)
}

// Find() => binary search implementation returns a pointer to the node and true if val is found in tree, else returns zero struct and false
func (t BSTree[T]) Find(val T) (*Node[T], bool) {
	zero := &Node[T]{}

	if t.Root == nil {
		return zero, false
	}

	for t.Root != nil {
		if val == t.Root.Data {
			return t.Root, true
		} else if val < t.Root.Data {
			t.Root = t.Root.Left
		} else if val > t.Root.Data {
			t.Root = t.Root.Right
		}
	}

	fmt.Println("Data not Found")
	return zero, false
}

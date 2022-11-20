package binarytree

import (
	"fmt"
	"io"
)

// Node struct
type Node struct {
	left  *Node
	right *Node
	Data  int64
}

// BTree struct
type BTree struct {
	Root *Node
}

// New instance builder
func New() *BTree {
	return &BTree{}
}

// Add element to the tree
func (t *BTree) Add(data int64) *BTree {
	if t.Root == nil {
		t.Root = &Node{Data: data, left: nil, right: nil}
	} else {
		t.Root.Add(data)
	}
	return t
}

// Add element to the node
func (n *Node) Add(data int64) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.left == nil {
			n.left = &Node{Data: data, left: nil, right: nil}
		} else {
			n.left.Add(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{Data: data, left: nil, right: nil}
		} else {
			n.right.Add(data)
		}
	}
}

// Dump btree value
func Dump(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	Dump(w, node.left, ns+2, 'L')
	Dump(w, node.right, ns+2, 'R')
}

package ch15binarysearchtree

import "fmt"

// TreeNode represents a binary tree structure
// contains value of the node and its left and right node
type TreeNode struct {
	Value 	int
	Left	*TreeNode
	Right	*TreeNode
}

// Insert inserts the node
func (t *TreeNode) Insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: v}
		} else {
			t.Left.Insert(v)
		}
	} else if v > t.Value {
		if t.Right == nil {
			t.Right = &TreeNode{Value: v}
		} else {
			t.Right.Insert(v)
		}
	}
}

// PrintInOrder prints the tree node in order
// left nodes, root node, and right nodes
func (t TreeNode) PrintInOrder() {
	if t.Left != nil {
		t.Left.PrintInOrder()
	}
	fmt.Printf("%v ", t.Value)
	if t.Right != nil {
		t.Right.PrintInOrder()
	}
}

func (t TreeNode) ToInOrderList() []int {
	var list []int
	list = t.toInOrderList(list)
	return list
}

func (t TreeNode) toInOrderList(list []int) []int{
	if t.Left != nil {
		list = t.Left.toInOrderList(list)
	}
	list = append(list, t.Value)
	if t.Right != nil {
		list = t.Right.toInOrderList(list)
	}
	return list
}

func (t TreeNode) ToPreOrderList() []int {
	list := []int{}
	list = t.toPreOrderList(list)
	return list
}

func (t TreeNode) toPreOrderList(list []int) []int {
	list = append(list, t.Value)
	if t.Left != nil {
		list = t.Left.toPreOrderList(list)
	}
	if t.Right != nil {
		list = t.Right.toPreOrderList(list)
	}
	return list
}

const (
	MinUint uint = 0                 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint      = ^MinUint          // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	MaxInt       = int(MaxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	MinInt       = ^MaxInt           // binary: all zeroes except high bit
)



package ch10tree

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: v}
		} else {
			t.Left.Insert(v)
		}
	} else {
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
	fmt.Println(t.Value)
	if t.Right != nil {
		t.Right.PrintInOrder()
	}
}

func (t TreeNode) ToInOrderList() []int {
	var list []int
	t.toInOrderList(list)
	return list
}

func (t TreeNode) toInOrderList(list []int) []int {
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
	var list []int
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

package ch15binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheFirstKeyGreaterThanGivenValueInBST1(t *testing.T) {
	/*
					10
				   /  \
				  5	  20
				 / \  / \
				4  7 15 25
		             /  /
		            12 22
	*/

	tree := &TreeNode{Value: 10}
	tree.Left = &TreeNode{Value: 5}
	tree.Left.Left = &TreeNode{Value: 4}
	tree.Left.Right = &TreeNode{Value: 7}

	tree.Right = &TreeNode{Value: 20}
	tree.Right.Left = &TreeNode{Value: 15}
	tree.Right.Left.Left = &TreeNode{Value: 12}
	tree.Right.Right = &TreeNode{Value: 25}
	tree.Right.Right.Left = &TreeNode{Value: 22}

	firstGreater := FindTheFirstKeyGreaterThanGivenValueInBST1(tree, 20)
	e := 22
	assert.Equal(t, e, firstGreater.Value)
}

func TestFindTheFirstKeyGreaterThanGivenValueInBST2(t *testing.T) {
	tree := &TreeNode{Value: 10}
	tree.Left = &TreeNode{Value: 5}
	tree.Left.Left = &TreeNode{Value: 4}
	tree.Left.Right = &TreeNode{Value: 7}

	tree.Right = &TreeNode{Value: 20}
	tree.Right.Left = &TreeNode{Value: 15}
	tree.Right.Left.Left = &TreeNode{Value: 12}
	tree.Right.Right = &TreeNode{Value: 25}
	tree.Right.Right.Left = &TreeNode{Value: 22}

	firstGreater := FindTheFirstKeyGreaterThanGivenValueInBST2(tree, 20)
	e := 22
	assert.Equal(t, e, firstGreater.Value)
}

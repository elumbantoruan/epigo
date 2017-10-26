package ch15binarysearchtree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsBST(t *testing.T) {
	tree := &TreeNode{Value:20}
	tree.Left = &TreeNode{Value:10}
	tree.Left.Left = &TreeNode{Value:5}
	tree.Right = &TreeNode{Value:30}

	b := IsBST(tree)
	e := true
	assert.Equal(t, e, b)

	tree = &TreeNode{Value:20}
	tree.Left = &TreeNode{Value:10}
	tree.Left.Left = &TreeNode{Value:15}
	tree.Right = &TreeNode{Value:30}

	b = IsBST(tree)
	e = false
	assert.Equal(t, e, b)
}

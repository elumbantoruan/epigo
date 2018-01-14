package ch10tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeNode_Insert(t *testing.T) {
	tree := &TreeNode{Value: 10}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(20)
	tree.Insert(15)

	rootVal := 10
	assert.True(t, tree.Value == rootVal)
	assert.True(t, tree.Left.Value < rootVal)
	assert.True(t, tree.Right.Value > rootVal)
}

func TestTreeNode_PrintInOrder(t *testing.T) {
	tree := &TreeNode{Value: 10}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(20)
	tree.Insert(15)

	tree.PrintInOrder()

	assert.True(t, tree != nil)
}

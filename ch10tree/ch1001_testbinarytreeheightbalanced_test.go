package ch10tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBinaryTreeHeightBalanced(t *testing.T) {
	tree := &TreeNode{Value: 20}
	tree.Left = &TreeNode{Value: 16}
	tree.Left.Left = &TreeNode{Value: 14}
	tree.Left.Left.Left = &TreeNode{Value: 10}

	tree.Right = &TreeNode{Value: 25}

	balanced := IsBinaryTreeHeightBalanced(tree)
	expected := false

	assert.Equal(t, expected, balanced)
}

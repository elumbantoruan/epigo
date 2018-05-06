package ch10tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBinaryTreeHeightImBalanced(t *testing.T) {
	tree := &TreeNode{Value: 20}
	tree.Left = &TreeNode{Value: 16}
	tree.Left.Left = &TreeNode{Value: 14}
	tree.Left.Left.Left = &TreeNode{Value: 10}

	tree.Right = &TreeNode{Value: 25}

	balanced := IsBinaryTreeHeightBalanced(tree)
	expected := false

	assert.Equal(t, expected, balanced)
}

func TestIsBinaryTreeHeightBalanced(t *testing.T) {
	tree := &TreeNode{Value: 20}
	tree.Left = &TreeNode{Value: 16}
	tree.Left.Left = &TreeNode{Value: 14}
	tree.Left.Left.Left = &TreeNode{Value: 10}

	tree.Right = &TreeNode{Value: 25}
	tree.Right.Right = &TreeNode{Value: 27}
	tree.Right.Right.Right = &TreeNode{Value: 37}

	balanced := IsBinaryTreeHeightBalanced(tree)
	expected := true

	assert.Equal(t, expected, balanced)
}

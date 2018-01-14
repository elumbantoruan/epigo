package ch15binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLCA(t *testing.T) {
	/*
					10
				   /  \
				  5	  20
				 / \  / \
				4  7 15 25
		             /
		            12
	*/

	tree := &TreeNode{Value: 10}
	tree.Left = &TreeNode{Value: 5}
	tree.Left.Left = &TreeNode{Value: 4}
	tree.Left.Right = &TreeNode{Value: 7}

	tree.Right = &TreeNode{Value: 20}
	tree.Right.Left = &TreeNode{Value: 15}
	tree.Right.Left.Left = &TreeNode{Value: 12}
	tree.Right.Right = &TreeNode{Value: 25}

	node := FindLCA(tree, 7, 25)
	e := 10
	assert.Equal(t, e, node.Value)

	node = FindLCA(tree, 12, 25)
	e = 20
	assert.Equal(t, e, node.Value)

}

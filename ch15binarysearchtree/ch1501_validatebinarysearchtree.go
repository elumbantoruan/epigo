package ch15binarysearchtree

import "math"

// IsBST tests if binary tree satisfies binary search tree condition
// all left nodes' value are less than root value
// and all right nodes' value are greater than root value
// Given below tree, IsBST is true
//						20
//					   /  \
//                    10  30
//                   /
//					5
// Given below tree, IsBST is false  (15 shouldn't be the child node of 10)
//						20
//					   /  \
//                    10  30
//                   /
//					15
//
func IsBST(tree *TreeNode) bool {
	return isBST(tree, math.MinInt32, math.MaxInt32)
}

func isBST(tree *TreeNode, minValue, maxValue int) bool {
	if tree == nil {
		return true
	}
	if tree.Value < minValue || tree.Value > maxValue {
		return false
	}
	return isBST(tree.Left, minValue, tree.Value) &&
		isBST(tree.Right, tree.Value, maxValue)
}
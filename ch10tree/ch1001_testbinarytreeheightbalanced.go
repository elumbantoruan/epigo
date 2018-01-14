package ch10tree

import "math"

func IsBinaryTreeHeightBalanced(node *TreeNode) bool {
	lh := height(node.Left)
	rh := height(node.Right)

	if math.Abs(float64(lh-rh)) > 1 {
		return false
	}

	return IsBinaryTreeHeightBalanced(node.Left) && IsBinaryTreeHeightBalanced(node.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
}

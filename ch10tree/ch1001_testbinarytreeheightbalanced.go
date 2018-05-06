package ch10tree

// IsBinaryTreeHeightBalanced .
func IsBinaryTreeHeightBalanced(node *TreeNode) bool {
	if node == nil {
		return true
	}
	lh := height(node.Left)
	rh := height(node.Right)

	// if abs(lh-rh) > 1 {
	// 	return false
	// }

	return abs(lh-rh) <= 1 && IsBinaryTreeHeightBalanced(node.Left) && IsBinaryTreeHeightBalanced(node.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

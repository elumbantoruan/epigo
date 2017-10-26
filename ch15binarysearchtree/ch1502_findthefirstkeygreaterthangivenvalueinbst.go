package ch15binarysearchtree

func FindTheFirstKeyGreaterThanGivenValueInBST1(tree *TreeNode, k int) *TreeNode {
	node := tree
	for node.Value != k {
		if node.Value > k {
			node = node.Left
		} else { //node.Value < k
			node = node.Right
		}
	}

	if node.Right != nil && node.Right.Left != nil {
		return node.Right.Left
	}
	if node.Right != nil && node.Right.Left == nil {
		return node.Right
	}
	return nil
}

func FindTheFirstKeyGreaterThanGivenValueInBST2(tree *TreeNode, k int) *TreeNode {
	node := tree
	firstGreater := &TreeNode{}

	for node != nil {
		if node.Value > k {
			firstGreater = node
			node = node.Left
		} else { //node.Value < k
			node = node.Right
		}
	}
	return firstGreater
}
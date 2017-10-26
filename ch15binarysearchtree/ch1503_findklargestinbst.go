package ch15binarysearchtree

// FindKLargestInBST returns k largest elements from the tree
// starts with the largest
// To do that, we have to backward by doing recursive routine which
// creates a stack so at the end we pop the element starting from the latest
// (which is the greatest) and start populating the list
// Time complexity is O(h+k) where h is the height of the nodes, and k is the number of elements to put into list
// Given below tree nodes and k = 3, the results will be {53,47,43}
//								19
//							   /  \
//							  /    \
//							 /      \
//                          7       43
//                         / \     / \
//                        3  11   23 47
//						 / \  \   \   \
//						2  5  17   37  53
//                            /   /  \
//                           13  29  41
//                                \
//                                31
//
func FindKLargestInBST(tree *TreeNode, k int) []int {
	list := []int{}
	findKLargestInBST(tree, k, list)
	return list
}

func findKLargestInBST(tree *TreeNode, k int, list []int) {
	if tree != nil && len(list) < k {
		findKLargestInBST(tree.Right, k, list)
		if len(list) < k {
			list = append(list, tree.Value)
			findKLargestInBST(tree.Left, k, list)
		}
	}
}

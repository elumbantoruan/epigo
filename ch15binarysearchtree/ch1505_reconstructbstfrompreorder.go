package ch15binarysearchtree

// ReconstructBSTFromPreOrder builds BST given a list of preOrder traversal data
// Given the preOrder list 19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47, 53
// it will construct the tree as follows:
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
// Time complexity is O(n)
func ReconstructBSTFromPreOrder(list []int) *TreeNode {
	return reconstructBSTFromPreOrder(list, 0, len(list))
}

func reconstructBSTFromPreOrder(list []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}
	transpoint := start + 1
	for transpoint < end {
		// advancing the index so it reaches the node before the end index
		if list[transpoint] < list[start] {
			transpoint++
		} else {
			break
		}
	}
	root := &TreeNode{Value: list[start]}
	root.Left = reconstructBSTFromPreOrder(list, start+1, transpoint)
	root.Right = reconstructBSTFromPreOrder(list, transpoint, end)

	return root
}

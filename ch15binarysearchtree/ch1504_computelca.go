package ch15binarysearchtree

// FindLCA returns a root to a and b
//
//			10
//		   /  \
//		  5	  20
//		 / \  / \
//		4  7 15 25
//			 /
//			12
// Given above tree nodes, with a = 7 and b = 25, the lca = 10
// a = 12 and b = 25, lca = 20
func FindLCA(root *TreeNode, a, b int) *TreeNode {
	current := root
	for current.Value < a || current.Value > b {
		for current.Value < a {
			current = current.Right
		}
		for current.Value > b {
			current = current.Left
		}
	}
	return current
}

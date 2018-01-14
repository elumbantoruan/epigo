package ch15binarysearchtree

// BuildBSTFromSortedList builds BST using
// similar approach as binary search.
// take mid value as the root, and smaller values will be located
// in the left nodes and reversely for greater values
// Time complexity T(n) = 2T(n/2) + O(1), which solves T(n) to O(n)
func BuildBSTFromSortedList(list []int) *TreeNode {
	return buildBSTFromSortedList(list, 0, len(list))
}

func buildBSTFromSortedList(list []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}

	mid := start + (end-start)/2
	tree := &TreeNode{Value: list[mid]}
	tree.Left = buildBSTFromSortedList(list, start, mid)
	tree.Right = buildBSTFromSortedList(list, mid+1, end)
	return tree
}

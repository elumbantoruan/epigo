package ch15binarysearchtree

type Interval struct {
	Left  int
	Right int
}

// RangeLookup return a list which values are in between the range of interval left and right
func RangeLookup(tree *TreeNode, interval Interval) []int {
	results := []int{}
	rangeLookup(tree, interval, results)
	return results
}

func rangeLookup(tree *TreeNode, interval Interval, results []int) {
	if tree == nil {
		return
	}

	if tree.Value >= interval.Left && tree.Value <= interval.Right {
		rangeLookup(tree.Left, interval, results)
		results = append(results, tree.Value)
		rangeLookup(tree.Right, interval, results)
	} else if tree.Value > interval.Right {
		rangeLookup(tree.Left, interval, results)
	} else {
		rangeLookup(tree.Right, interval, results)
	}
}

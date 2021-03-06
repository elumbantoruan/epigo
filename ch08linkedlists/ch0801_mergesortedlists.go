package ch08linkedlists

// MergeTwoSortedList merges two sorted lists
// newNode is created for placeholder for a result
// Time complexity is O(n+m) where n and m are for l1 and l2 respectively
// In the best case, one list is much shorter
// The space complexity is O(1) since we use the existing nodes
func MergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	current := newNode
	pl1 := l1
	pl2 := l2
	for pl1 != nil && pl2 != nil {
		if pl1.Value.(int) < pl2.Value.(int) {
			current.Next = pl1
			pl1 = pl1.Next
		} else {
			current.Next = pl2
			pl2 = pl2.Next
		}
		current = current.Next
	}
	if pl1 == nil {
		current.Next = pl2
	} else {
		current.Next = pl1
	}

	return newNode.Next
}

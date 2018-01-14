package ch08linkedlists

// OverlappingLists identifies the overlap node between the two nodes
// Time complexity is O(n)
// Space complexity is O(1)
func OverlappingLists(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := l1.Length()
	len2 := l2.Length()

	if len1 > len2 {
		//l1 = advanceNode(l1, len1 - len2)
		l1 = l1.AdvanceNode(len1 - len2)
	} else {
		//l2 = advanceNode(l2, len2 - len1)
		l2 = l2.AdvanceNode(len2 - len1)
	}
	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return nil
}

func advanceNode(l *ListNode, n int) *ListNode {
	i := 0
	for i < n {
		l = l.Next
		i++
	}
	return l
}

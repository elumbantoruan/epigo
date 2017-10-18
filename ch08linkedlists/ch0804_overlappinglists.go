package ch08linkedlists


func OverlappingLists (l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := listLength(l1)
	len2 := listLength(l2)

	if len1 > len2 {
		l1 = advanceNode(l1, len1 - len2)
	} else {
		l2 = advanceNode(l2, len1 - len2)
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

func listLength(l *ListNode) int {
	i := 0
	c := l
	for c != nil {
		c = c.Next
		i++
	}
	return i
}

func advanceNode(l *ListNode, n int) *ListNode {
	i := 0
	for i < n {
		l = l.Next
		i++
	}
	return l
}

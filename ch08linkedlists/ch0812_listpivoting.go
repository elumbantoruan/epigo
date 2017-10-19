package ch08linkedlists

// ListPivoting pivots for a given k
// Input list: { 3, 2, 2, 11, 7, 5, 11 }
// k: 7
// all nodes with value less than 7 should be located prior to 7
// and all nodes with value greater than 7 should be located after
// Output:     { 3, 2, 2, 5, 7, 11, 11 }
// Time complexity involves three lists which rounds to O(n)
// Space complexity is O(1)
func ListPivoting(l *ListNode, k int) *ListNode {
	lessHead := &ListNode{}
	equalHead := &ListNode{}
	greaterHead := &ListNode{}

	lessIterator := lessHead
	equalIterator := equalHead
	greaterIterator := greaterHead

	p := l

	for p != nil {
		if p.Value.(int) < k {
			lessIterator.Next = p
			lessIterator = lessIterator.Next
		} else if p.Value.(int) == k {
			equalIterator.Next = p
			equalIterator = equalIterator.Next
		} else {
			greaterIterator.Next = p
			greaterIterator = greaterIterator.Next
		}
		p = p.Next
	}

	greaterIterator.Next = nil
	equalIterator.Next = greaterHead.Next
	lessIterator.Next = equalHead.Next

	return lessHead.Next
}

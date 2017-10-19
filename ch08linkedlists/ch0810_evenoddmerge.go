package ch08linkedlists

// EvenOddMerge merges the nodes into a sequence of even numbers followed by odd numbers
// Time complexity is O(n)
// Space complexity is O(1)
func EvenOddMerge(l *ListNode) *ListNode {

	p1 := l
	p2 := l.Next
	connect := l.Next

	for p1 != nil && p2 != nil && p2.Next != nil {

		p1.Next = p2.Next
		p1 = p1.Next

		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = connect
	return l
}

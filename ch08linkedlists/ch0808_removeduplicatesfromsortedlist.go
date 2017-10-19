package ch08linkedlists

// RemoveDuplicatesFromSortedList removes the duplicate node
// Given 	{ 2, 2, 3, 5, 7, 11, 11 }
// output 	{ 2, 3, 5, 7, 11 }
// Time complexity is O(n)
// Space complexity is O(1)
func RemoveDuplicatesFromSortedList(l *ListNode) *ListNode {
	p := l
	for p.Next != nil {
		if p.Value == p.Next.Value {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return l
}

package ch08linkedlists

// RemoveKthLastElement removes the kth last element from linkedlist
// Given list of { 1, 2, 3, 4, 5, 6, 7, 8 }
// and k = 4
// The output will be { 1, 2, 3, 4, 6, 7, 8 }
// Time complexity is O(n)
// Space complexity is O(1)
func RemoveKthLastElement(l *ListNode, k int) *ListNode {
	newNode := &ListNode{}
	newNode.Next = l
	rear := l
	front := l

	for k > 0 {
		front = front.Next
		k--
	}

	for front != nil {
		rear = rear.Next
		front = front.Next
	}

	rear.Value = rear.Next.Value
	rear.Next = rear.Next.Next

	return newNode.Next
}

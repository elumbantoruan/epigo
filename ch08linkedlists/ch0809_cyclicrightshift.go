package ch08linkedlists

// CyclicRightShift shifts the nodes from Kth to the right
// Given:   { 2, 3, 5, 3, 2 }
// and k: 3
// Output:	{ 5, 3, 2, 2, 3 }
// NOTE: This solution uses a cyclic list to make it more efficient
//       rather than creating a new node by copying its value
// First calculate the length, and find out the nodes to be shifted.
// k can be greater than length, or it can be the same
// do k%length to find out how may nodes to be shifted
// Time complexity O(n) and space complexity is O(1)
func CyclicRightShift(l *ListNode, k int) *ListNode {
	// create tail as an iterator
	// length starts with 1 and iterates with tail.Next
	// so we still get the end of the node with not nil reference
	tail := l
	len := 1
	for tail.Next != nil {
		tail = tail.Next
		len++
	}
	// get mod k with len
	k = k % len
	if k == 0 {
		return l
	}

	// now tail points to last node and connect the next node with l
	// to make a list cycle
	tail.Next = l

	// need to shift the node
	shift := len - k
	newTail := tail
	for shift > 0 {
		newTail = newTail.Next
		shift--
	}
	// newHead points to newTail.Next as a new head of list
	newHead := newTail.Next
	// but this newTail is cyclic so we need to cut it off
	newTail.Next = nil

	return newHead

}

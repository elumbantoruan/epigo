package ch08linkedlists

// ReverseSubList reverses the nodes from the start and finish index of nodes
// The index starts with 1
//   11 3 5 7 2
//   11 7 5 3 2
func ReverseSubList(l *ListNode, start, finish int) *ListNode {
	idx := 1
	newNode := &ListNode{}
	newNode.Next = l
	current := newNode

	for idx < start {
		current = current.Next
		idx++
	}

	sublist := current.Next

	for start < finish {
		next := sublist.Next
		sublist.Next = next.Next
		next.Next = current.Next
		current.Next = next
		start++
	}
	return newNode.Next
}

package ch08linkedlists

// ReverseSubList reverses the nodes from the start and finish index of nodes
// The index starts with 1
// input:   11 3 5 7 2
// start:  2
// finish: 4
// output:  11 7 5 3 2
// NOTE: No new node should be allocated
// Time complexity is 0(2n).  One is counted for slow node and the other one is for fast node.
// Space complexity is 0(1)
func ReverseSubList(l *ListNode, start, finish int) *ListNode {
	idx := 1
	newNode := &ListNode{}
	newNode.Next = l
	current := newNode

	for idx < start {
		current = current.Next
		idx++
	}

	// we identify the sublist node
	sublist := current.Next

	// perform the update
	// doing this it will not require allocation of new node
	for start < finish {
		next := sublist.Next
		sublist.Next = next.Next
		next.Next = current.Next
		current.Next = next
		start++
	}
	return newNode.Next
}

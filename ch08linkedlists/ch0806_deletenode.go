package ch08linkedlists

// DeleteNode deletes a node in single linked list
// This approach is only applicable for node that's not the end of the list
// Time complexity is O(1)
func DeleteNode(node *ListNode) {
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

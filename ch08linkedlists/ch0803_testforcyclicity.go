package ch08linkedlists

// TestForCyclicity checks if there is a cycle in ListNode
// Let F be the number of nodes to the start of the cycle
// C the number of nodes on the cycle
// n the total number of nodes
// The time complexity is O(F) + O(C) = O(n) - O(F)
// The space complexity is O(1)
func TestForCyclicity(l *ListNode) *ListNode {
	slow := l
	fast := l

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}
	if fast == nil {
		return nil
	}
	slow = l
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

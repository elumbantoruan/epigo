package ch08linkedlists

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

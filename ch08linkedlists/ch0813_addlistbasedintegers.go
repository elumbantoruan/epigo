package ch08linkedlists

// AddListBasedIntegers adds two ListNode
// Time complexity is O(n+m)
// Space complexity is O(max(n,m))
// where n is l1 and m is l2
func AddListBasedIntegers(l1, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	ni := newNode

	carryOver := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Value.(int)
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Value.(int)
			l2 = l2.Next
		}
		sum += carryOver
		carryOver = sum / 10
		ni.Next = &ListNode{Value:sum % 10}
		ni = ni.Next
	}

	if carryOver > 0 {
		ni.Next = &ListNode{Value:carryOver}
		ni = ni.Next
	}

	return newNode.Next
}

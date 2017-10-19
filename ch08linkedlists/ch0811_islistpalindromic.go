package ch08linkedlists

// IsPalindromic tests the node if its values are palindrome
// Time complexity is O(n)
// Space complexity is O(1)
func IsPalindromic(l *ListNode) bool {
	slow := l
	fast := l

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := Reverse(slow)
	fistHalf := l

	for secondHalf != nil {
		if fistHalf.Value != secondHalf.Value {
			return false
		}
		fistHalf = fistHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}

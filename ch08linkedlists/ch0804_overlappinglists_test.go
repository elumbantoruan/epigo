package ch08linkedlists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverlappingLists(t *testing.T) {

	l1 := &ListNode{Value: 4}
	l1.Next = &ListNode{Value: 5}
	l1.Next.Next = &ListNode{Value: 6}
	l1.Next.Next.Next = &ListNode{Value: 7}
	l1.Next.Next.Next.Next = &ListNode{Value: 8}

	l2 := &ListNode{Value: 1}
	l2.Next = l1.Next.Next

	ovr := OverlappingLists(l1, l2)

	assert.True(t, ovr != nil, "should have an overlap")

	expected := 6
	assert.Equal(t, expected, ovr.Value.(int), "doesn't match")
}

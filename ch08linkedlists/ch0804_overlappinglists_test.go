package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOverlappingLists(t *testing.T) {
	l1 := &ListNode{Value:1}
	l1.Next = &ListNode{Value:2}

	l2 := &ListNode{Value:8}
	l2.Next = l1.Next
	l2.Next.Next = &ListNode{Value:4}
	l2.Next.Next.Next = &ListNode{Value:5}

	ovr := OverlappingLists(l1, l2)

	assert.True(t, ovr != nil, "should have an overlap")

	expected := 2
	assert.Equal(t, expected, ovr.Value.(int), "doesn't match")
}

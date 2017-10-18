package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTestForCyclicity(t *testing.T) {
	l := &ListNode{Value:1}
	l.Next = &ListNode{Value:2}
	l.Next.Next = &ListNode{Value:3}
	l.Next.Next.Next = &ListNode{Value:4}
	l.Next.Next.Next.Next = &ListNode{Value:5}
	l.Next.Next.Next.Next.Next = l.Next

	c := TestForCyclicity(l)
	assert.True(t, c != nil, "should be cyclic")

	expected := 2
	assert.Equal(t, expected, c.Value.(int), "should match")

}

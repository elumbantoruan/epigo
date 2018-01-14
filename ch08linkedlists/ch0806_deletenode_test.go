package ch08linkedlists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node := &ListNode{Value: 1}
	node.Next = &ListNode{Value: 2}
	node.Next.Next = &ListNode{Value: 3}
	node.Next.Next.Next = &ListNode{Value: 4}

	DeleteNode(node.Next.Next)

	actual := node.ToArray()

	expected := []interface{}{1, 2, 4}

	assert.True(t, actual != nil, "result is nil")

	for i, e := range actual {
		assert.Equal(t, expected[i], e, "not matched")
	}
}

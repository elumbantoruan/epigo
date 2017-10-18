package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

// TestReverseSubList reverses the nodes from the start and finish index of nodes
// The index starts with 1
//   11 3 5 7 2
//   11 7 5 3 2
func TestReverseSubList(t *testing.T) {
	node := new(ListNode)
	node.Value = 11
	node.Next = &ListNode{Value:3}
	node.Next.Next = &ListNode{Value:5}
	node.Next.Next.Next = &ListNode{Value:7}
	node.Next.Next.Next.Next = &ListNode{Value:2}

	start := 2
	finish := 4

	fmt.Printf("\nTestReverseSubList input.  Start: %d Finish: %d\n", start, finish)
	node.Print()

	reversed := ReverseSubList(node, start, finish)

	fmt.Println("TestReverseSubList output:")
	reversed.Print()

	expected := []interface{} {11, 7, 5, 3, 2}
	actual := reversed.ToArray()

	assert.True(t, len(expected) == len(actual))

	for i, e := range actual {
		assert.Equal(t, expected[i], e, "not matched")
	}

}
package ch08linkedlists


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestMergeTwoSortedList(t *testing.T) {
	l1 := &ListNode{Value:2}
	l1.Next = &ListNode{Value:5}
	l1.Next.Next = &ListNode{Value:7}

	l2 := &ListNode{Value:3}
	l2.Next = &ListNode{Value:11}

	result := MergeTwoSortedList(l1, l2)

	assert.True(t, result != nil)

	arr := result.ToArray()

	expected := []interface{} { 2,3,5,7,11 }
	for i, e := range arr {
		fmt.Printf("%v ", e)
		assert.Equal(t, expected[i], e, "not match")
	}
}

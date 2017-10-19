package ch08linkedlists

import (
	"fmt"
)

// ListNode represents a single linked list
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// Print prints the ListNode
func (l *ListNode) Print() {
	p := l
	for p != nil {
		fmt.Printf("%v ", p.Value)
		p = p.Next
	}
	fmt.Println()
}

// ToArray converts the nodes of ListNode into an array (interface{})
func (l *ListNode) ToArray() ([]interface{}) {
	var values []interface{}
	p := l
	for p != nil {
		values = append(values, p.Value)
		p = p.Next
	}
	return values
}

// Length returns the length of ListNode
func (l *ListNode) Length() int {
	c := 0
	p := l
	for p != nil {
		p = p.Next
		c++
	}
	return c
}

// AdvanceNode advances the node for n steps
func (l *ListNode) AdvanceNode(n int) *ListNode {
	p := l
	for n > 0 {
		p = p.Next
		n--
	}
	return p
}

// FromArray creates a ListNode from a given array
func FromArray(arr []interface{}) (*ListNode) {
	node := &ListNode{}
	p := node
	for _, e := range arr {
		p.Next = &ListNode{Value:e}
		p = p.Next
	}
	return node.Next
}

// AdvanceNode advances the node for n steps
func AdvanceNode(l *ListNode, n int) *ListNode{
	for n > 0 {
		l = l.Next
		n--
	}
	return l
}

// Reverse reverses the nodes
func Reverse(l *ListNode) *ListNode {
	c := l
	var prev *ListNode

	for c != nil {
		next := c.Next
		c.Next = prev

		prev = c
		c = next
	}
	return prev
}

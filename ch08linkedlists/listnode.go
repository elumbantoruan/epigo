package ch08linkedlists

import (
	"fmt"
)

// ListNode represents a single linked list
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// Print prints the list node
func (l *ListNode) Print() {
	p := l
	for p != nil {
		fmt.Printf("%v ", p.Value)
		p = p.Next
	}
	fmt.Println()
}

func (l *ListNode) ToArray() ([]interface{}) {
	var values []interface{}
	p := l
	for p != nil {
		values = append(values, p.Value)
		p = p.Next
	}
	return values
}
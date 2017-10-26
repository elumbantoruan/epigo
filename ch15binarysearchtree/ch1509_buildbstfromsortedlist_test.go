package ch15binarysearchtree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuildBSTFromSortedList(t *testing.T) {
	list := []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 }
	tree := BuildBSTFromSortedList(list)
	inordertraversal := tree.ToInOrderList()
	for i, e := range inordertraversal {
		assert.Equal(t, list[i], e)
	}
}

package ch15binarysearchtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeLookup(t *testing.T) {
	tree := &TreeNode{Value: 19}
	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(11)
	tree.Insert(17)
	tree.Insert(13)

	tree.Insert(43)
	tree.Insert(23)
	tree.Insert(37)
	tree.Insert(29)
	tree.Insert(41)
	tree.Insert(31)
	tree.Insert(47)
	tree.Insert(53)

	results := RangeLookup(tree, Interval{3, 17})
	expected := []int{3, 5, 7, 11, 13, 17}
	for i, r := range results {
		assert.Equal(t, expected[i], r)
	}
}

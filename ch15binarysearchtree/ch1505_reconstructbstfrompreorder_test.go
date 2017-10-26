package ch15binarysearchtree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReconstructBSTFromPreOrder(t *testing.T) {
	preOrderList := []int {19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47, 53}
	tree := ReconstructBSTFromPreOrder(preOrderList)
	expected := tree.ToPreOrderList()
	for i, p := range preOrderList {
		assert.Equal(t, expected[i], p)
	}
}

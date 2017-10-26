package ch14sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergingTwoSortedArrays(t *testing.T) {
	l1 := []int{ 5, 13, 17, 0, 0, 0, 0, 0 }
	m := 3
	l2 := []int{ 3, 7, 11, 19 }
	n := 4
	MergingTwoSortedArrays(l1, m, l2, n)
	expected := []int {3, 5, 7, 11, 13, 17, 19, 0}

	for i, v := range l1 {
		assert.Equal(t, expected[i], v)
	}
}

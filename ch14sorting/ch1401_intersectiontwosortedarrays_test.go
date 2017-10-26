package ch14sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntersectionTwoSortedArrays(t *testing.T) {
	l1 := []int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}
	l2 := []int{5, 5, 6, 8, 8, 9, 10, 10}
	ints := IntersectionTwoSortedArrays(l1, l2)
	expected := []int{5,6,8}
	for i, v := range ints {
		assert.Equal(t, expected[i], v)
	}
}

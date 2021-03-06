package ch12searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//{-2, 0, 2, 3, 6, 7, 9}
func TestSearchSortedArrayForEntryEqualToIndex(t *testing.T) {
	numbers := []int{-2, 0, 2, 3, 6, 7, 9}
	result := SearchSortedArrayForEntryEqualToIndex(numbers)
	expected := []int{2, 3}
	assert.Contains(t, expected, result)
}

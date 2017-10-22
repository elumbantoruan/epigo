package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicatesFromSortedArray(t *testing.T) {
	numbers := []int {2,3,5,5,7,11,11,11,13}
	// after deletion, the array is {2,3,5,7,11,13,0,0,0}
	result := DeleteDuplicatesFromSortedArray(numbers)
	expected := 6
	assert.Equal(t, expected, result)
}
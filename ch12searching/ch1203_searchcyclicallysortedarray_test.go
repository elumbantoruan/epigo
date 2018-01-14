package ch12searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchCyclicallySortedArray(t *testing.T) {
	numbers := []int{378, 478, 550, 631, 103, 203, 220, 234, 279, 368}
	index := SearchCyclicallySortedArray(numbers)
	expected := 4
	assert.Equal(t, expected, index)
}

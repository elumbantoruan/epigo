package ch12searching

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearchForFirstOccurenceOfK(t *testing.T) {
	numbers := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
	k := 285
	expected := 6
	n := SearchForFirstOccurenceOfK(numbers, k)
	assert.Equal(t, expected, n)
}

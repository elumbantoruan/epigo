package ch06arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	i := []int{9, 9}
	r := PlusOne(i)
	expected := []int{1, 0, 0}
	for i, n := range r {
		assert.Equal(t, expected[i], n)
	}

	i = []int{9, 8}
	r = PlusOne(i)
	expected = []int{9, 9}
	for i, n := range r {
		assert.Equal(t, expected[i], n)
	}
}

package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	i := []int { 9, 9}
	r := PlusOne(i)
	expected := []int{1,0,0}
	for i, n := range r {
		assert.Equal(t, expected[i], n)
	}
}

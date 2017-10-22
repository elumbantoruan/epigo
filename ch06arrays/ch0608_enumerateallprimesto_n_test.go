package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrimes(t *testing.T) {
	n := 18
	results := GeneratePrimes(n)
	expected := []int {2,3,5,7,11,13,17}
	for i, r := range results {
		assert.Equal(t, expected[i], r)
	}
}

package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// ReplaceAndRemove replace each 'a' by two 'd'
// and delete each entry containing a 'b
// input of {a, c, d, b, b, c, a}
// output   {d, d, c, d, c, d, d}
func TestReplaceAndRemove(t *testing.T) {
	runes := []rune{'a', 'c', 'd', 'b', 'b', 'c', 'a'}
	ReplaceAndRemove(runes)

	expected := []rune {'d', 'd', 'c', 'd', 'c', 'd', 'd'}
	for i:=0; i < len(runes); i++ {
		assert.Equal(t, expected[i], runes[i])
	}

}

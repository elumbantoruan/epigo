package ch08linkedlists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromic_True(t *testing.T) {
	l := FromArray([]interface{}{1, 2, 3, 4, 3, 2, 1})
	ok := IsPalindromic(l)
	assert.Equal(t, true, ok)
}

func TestIsPalindromic_False(t *testing.T) {
	l := FromArray([]interface{}{1, 2, 3, 4, 3, 6, 1})
	ok := IsPalindromic(l)
	assert.Equal(t, false, ok)
}

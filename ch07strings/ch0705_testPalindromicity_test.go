package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStringPalindrome(t *testing.T) {
	s := "A man, a plan, a canal, Panama"
	b := IsStringPalindrome(s)
	assert.Equal(t, true, b)

	s = "Ray a Ray"
	b = IsStringPalindrome(s)
	assert.Equal(t, false, b)
}

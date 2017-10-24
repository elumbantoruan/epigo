package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsStringPalindrome(t *testing.T) {
	s := "A man, a plan, a canal, Panama"
	b := IsStringPalindrome(s)
	assert.Equal(t, true, b)

	s = "Ray a Ray"
	b = IsStringPalindrome(s)
	assert.Equal(t, false, b)
}
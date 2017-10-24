package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Encoding encodes repeated characters
// by the repetitive count and character
// "aaaabcccaa" --> "4a1b3c2a"
func TestEncoding(t *testing.T) {
	s := "aaaabcccaa"
	r := Encoding(s)
	e := "4a1b3c2a"
	assert.Equal(t, e, r)
}

// Decoding decodes the number (count) into a repeated characters
// "3e4f2e" --> "eeeffffee"
func TestDecoding(t *testing.T) {
	s := "3e4f2e"
	e := "eeeffffee"
	r := Decoding(s)
	assert.Equal(t, e, r)
}

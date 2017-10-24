package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	s := "Hello World!"
	runes := []rune(s)
	ReverseWords(runes)
	reversed := "World! Hello"
	assert.Equal(t, reversed, string(runes))

}

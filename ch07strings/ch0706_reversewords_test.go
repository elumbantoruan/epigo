package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "Hello World!"
	runes := []rune(s)
	ReverseWords(runes)
	reversed := "World! Hello"
	assert.Equal(t, reversed, string(runes))

}

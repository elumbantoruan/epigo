package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearchSubstring(t *testing.T) {
	text := "this is a substring test"
	s := "substring"
	i := SearchSubstring(text, s)
	e := 10

	assert.Equal(t, e, i)

	text = "this is a test for substring"
	s = "substring"
	i = SearchSubstring(text, s)
	e = 19

	assert.Equal(t, e, i)
}
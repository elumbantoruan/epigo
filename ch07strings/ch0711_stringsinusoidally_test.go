package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringSinusoidally(t *testing.T) {
	s := "Hello World!"
	r := StringSinusoidally(s)
	e := "e lHloWrdlo!"

	assert.Equal(t, e, r)
}

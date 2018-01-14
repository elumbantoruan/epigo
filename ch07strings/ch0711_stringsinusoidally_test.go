package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSinusoidally(t *testing.T) {
	s := "Hello World!"
	r := StringSinusoidally(s)
	e := "e lHloWrdlo!"

	assert.Equal(t, e, r)
}

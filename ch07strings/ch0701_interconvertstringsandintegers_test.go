package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntegerToString(t *testing.T) {
	n := 123
	s := IntegerToString(n)
	expected := "123"
	assert.Equal(t, expected, s)
}

func TestStringToInteger(t *testing.T) {
	s := "123"
	n := StringToInteger(s)
	expected := 123
	assert.Equal(t, expected, n)
}

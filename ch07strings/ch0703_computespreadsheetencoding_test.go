package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpreadsheetEncoding(t *testing.T) {
	s := "D"
	n := SpreadsheetEncoding(s)
	expected := 4
	assert.Equal(t, expected, n)

	s = "AA"
	n = SpreadsheetEncoding(s)
	expected = 27
	assert.Equal(t, expected, n)

	s = "ZZ"
	n = SpreadsheetEncoding(s)
	expected = 702
	assert.Equal(t, expected, n)
}

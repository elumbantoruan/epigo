package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
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

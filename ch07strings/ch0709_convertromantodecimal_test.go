package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertRomanToDecimal(t *testing.T) {
	s := "XI"
	v := ConvertRomanToDecimal(s)
	e := 11
	assert.Equal(t, e, v)

	s = "IX"
	v = ConvertRomanToDecimal(s)
	e = 9
	assert.Equal(t, e, v)
}

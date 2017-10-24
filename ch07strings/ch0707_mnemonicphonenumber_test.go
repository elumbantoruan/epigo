package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMnemonicPhoneNumber(t *testing.T) {
	s := "2276696"
	mnemonics := MnemonicPhoneNumber(s)
	contains := false
	for _, c := range mnemonics {
		if c == "ACRONYM" {
			contains = true
			break
		}
	}
	assert.Equal(t, true, contains)
}

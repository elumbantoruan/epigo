package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

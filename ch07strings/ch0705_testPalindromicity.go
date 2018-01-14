package ch07strings

import (
	"strings"
	"unicode"
)

func IsStringPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	s = strings.ToLower(s)
	runes := []rune(s)

	for start < end {
		if unicode.IsLetter(runes[start]) && unicode.IsLetter(runes[end]) && runes[start] == runes[end] {
			start++
			end--
		} else if !unicode.IsLetter(runes[start]) && unicode.IsLetter(runes[end]) {
			start++
		} else if unicode.IsLetter(runes[start]) && !unicode.IsLetter(runes[end]) {
			end--
		} else {
			return false
		}

	}
	return true
}

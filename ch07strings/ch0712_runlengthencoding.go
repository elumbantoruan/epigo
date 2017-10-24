package ch07strings

import "unicode"

// Encoding encodes repeated characters
// by the repetitive count and character
// "aaaabcccaa" --> "4a1b3c2a"
func Encoding(s string) string {
	result := []rune{}
	count := 1
	runes := []rune(s)
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i -1] {
			count++
		} else if runes[i] != runes[i-1] {
			result = append(result, rune(count + '0'))
			result = append(result, runes[i - 1])
			count = 1
		}
		if i == len(runes) - 1{
			result = append(result, rune(count + '0'))
			result = append(result, runes[i])
		}
	}
	return string(result)
}

// Decoding decodes the number (count) into a repeated characters
// "3e4f2e" --> "eeeffffee"
func Decoding(s string) string {
	result := []rune{}
	runes := []rune(s)
	val := 0
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			val = int(runes[i] - '0')
		} else {
			for v := 0; v < val; v++ {
				result = append(result, runes[i])
			}
		}
	}
	return string(result)
}


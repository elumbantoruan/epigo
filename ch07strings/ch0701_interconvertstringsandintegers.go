package ch07strings

import (
	"math"
)

func IntegerToString(n int) string {
	len := int(math.Floor(math.Log10(float64(n)) + 1))
	var chars []rune
	for i := 0; i < len; i++ {
		chars = append(chars, ' ')
	}
	idx := len - 1
	for n > 0 {
		chars[idx] = rune(n%10 + '0')
		n /= 10
		idx--
	}

	return string(chars)
}

func StringToInteger(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = result*10 + int(s[i]-'0')
	}
	return result
}

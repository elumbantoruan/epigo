package ch05primitivetypes

import (
	"math"
)

// IsPalindrome checks if the reversed digits is palindrome
// for example: 323 is true while 1234 is false
func IsPalindrome(n int) bool {
	numDigits := int(math.Log10(float64(n))) + 1
	for n > 0 {
		least := n % 10
		most := n / int(math.Pow10(numDigits-1))
		if most != least {
			return false
		}
		n /= int(math.Pow10(numDigits - 1))
		n /= 10

		numDigits--
	}
	return true
}

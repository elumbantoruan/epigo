package ch05primitivetypes

import (
	"math"
)

func IsPalindrome(n int) bool {
	for n > 0 {
		numDigits := int(math.Floor(math.Log10(float64(n))) + 1)
		most := int( math.Ceil(float64( n / int(math.Pow10(int(numDigits - 1))))))
		least := n % 10
		if most != least {
			return false
		}
		n = n - most * int(math.Pow10(int(numDigits - 1)))
		n /= 10
	}
	return true
}

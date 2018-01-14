package ch05primitivetypes

// ReverseDigit returns a reversed digit
func ReverseDigit(n int) int {
	result := 0
	for n > 0 {
		result = result*10 + (n % 10)
		n /= 10
	}
	return result
}

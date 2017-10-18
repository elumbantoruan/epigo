package ch05primitivetypes

func ReverseDigit(n int) int {
	var result = 0
	for n > 0 {
		result = result * 10 + n % 10
		n /= 10
	}
	return  result
}

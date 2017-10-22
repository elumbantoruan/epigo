package ch06arrays

import "math"

func Multiply(n1, n2 []int) []int {
	sign := 1
	if n1[0] * n2[0] < 0 {
		sign = -1
	}

	n1[0] = int(math.Abs(float64(n1[0])))
	n2[0] = int(math.Abs(float64(n2[0])))

	r := make([]int, len(n1) + len(n2))
	carryOver := 0
	for i := len(n1) - 1; i >= 0 ; i-- {
		for j := len(n2) - 1; j >=0 ; j-- {
			val := n1[i] * n2[j] + carryOver
			r[i + j + 1] += val % 10
			carryOver = val / 10

			if carryOver > 0 {
				r[i + j] += carryOver
				carryOver = 0
			}
		}
	}
	if r[0] == 0 {
		r[1] *= sign
		return r[1:]
	}
	r[0] *= sign
	return r
}
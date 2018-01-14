package ch06arrays

// PlusOne adds one to the values of integer array
// For example: []int { 9, 8 } will be { 9, 9 }
//				[]int { 9, 9 } will be { 1, 0, 0 }
func PlusOne(i []int) []int {
	i[len(i)-1] += 1
	if i[len(i)-1] < 9 {
		return i
	}

	carryOver := 0
	for n := len(i) - 1; n >= 0; n-- {
		sum := carryOver + i[n]
		if sum > 9 {
			i[n] = 0
			carryOver = 1
		} else {
			i[n] = sum
			carryOver = 0
		}
	}

	if carryOver == 1 {
		r := make([]int, len(i)+1)
		r[0] = 1
		copy(r[1:], i)
		return r
	}

	return i
}

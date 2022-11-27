package ch06arrays

// PlusOne adds one to the values of integer array
// For example: []int { 9, 8 } will be { 9, 9 }
//
//	[]int { 9, 9 } will be { 1, 0, 0 }
func PlusOne(nums []int) []int {
	n := len(nums) - 1
	nums[n] += 1
	if nums[n] < 10 {
		return nums
	}

	carryOver := 0
	for n := len(nums) - 1; n >= 0; n-- {
		sum := carryOver + nums[n]
		if sum > 9 {
			nums[n] = 0
			carryOver = 1
		} else {
			nums[n] = sum
			carryOver = 0
		}
	}

	if carryOver == 1 {
		nums = append([]int{1}, nums...)
	}

	return nums
}

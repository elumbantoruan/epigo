package ch12searching

// SearchForFirstOccurenceOfK search first occurence of K in the list
// Given {-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
// and k 285
// The result is 6
func SearchForFirstOccurenceOfK(numbers []int, k int) int {
	index := -1
	floor := 0
	ceiling := len(numbers) - 1
	for floor <= ceiling {
		mid := floor + (ceiling-floor)/2
		if numbers[mid] == k {
			index = mid
			ceiling = mid - 1
		} else if numbers[mid] > k {
			ceiling = mid
		} else {
			floor = mid + 1
		}
	}
	return index
}

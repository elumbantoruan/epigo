package ch12searching

// SearchSortedArrayForEntryEqualToIndex returns an index of array
// which value is the same of the index
// Given numbers = {-2, 0, 2, 3, 6, 7, 9};
// index 2 or 3
func SearchSortedArrayForEntryEqualToIndex(numbers []int) int {
	floor := 0
	ceiling := len(numbers) - 1

	for floor <= ceiling {
		mid := floor + (ceiling-floor)/2
		if numbers[mid] == mid {
			return mid
		} else if numbers[mid] > mid {
			ceiling = mid
		} else {
			floor = mid + 1
		}
	}
	return -1
}

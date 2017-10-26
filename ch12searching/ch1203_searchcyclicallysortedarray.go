package ch12searching

 // SearchCyclicallySortedArrays search the smallest element in the array
 // and returns its index
 // Given numbers {378, 478, 550, 631, 103, 203, 220, 234, 279, 368};
 // The smallest element is 103, thus its index is 4
 func SearchCyclicallySortedArray(numbers []int) int {
 	floor := 0
 	ceiling := len(numbers) - 1
 	for floor <= ceiling {
		mid := floor + (ceiling - floor) / 2
		if numbers[mid] < numbers[floor] {
			floor = mid
		} else {
			ceiling = mid
		}
		if floor + 1 == ceiling {
			break
		}
	}
	return floor
 }

package ch06arrays

// DeleteDuplicatesFromSortedArray deletes duplicate array element
// given {2,3,5,5,7,11,11,11,13}
// after deletion, the array is {2,3,5,7,11,13,0,0,0}
// so the result is 6 (the number of valid entries)
func DeleteDuplicatesFromSortedArray(numbers []int) int {
	writeIdx := 1
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] != numbers[i] {
			numbers[writeIdx] = numbers[i]
			writeIdx++
		}
	}

	for i := writeIdx; i < len(numbers); i++ {
		numbers[i] = 0
	}
	return writeIdx
}

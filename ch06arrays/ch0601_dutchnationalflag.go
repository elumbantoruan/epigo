package ch06arrays

// DutchNationalFlag sorts the array given a pivorIndex so the element which values
// is the same as value of the pivotIndex stays together
//
// Keep the following invariants during partitioning
// bottom goup; A sublist(0, smaller)
// middle group: A sublist (smaller, equal)
// unclassified: A sublist (equal, larger)
// top: A sublist(larger, A.Count)
func DutchNationalFlag(colors []Color, pivotIndex int) {
	smaller := 0
	equal := 0
	larger := len(colors) -1

	pivot := colors[pivotIndex]

	for equal <= larger {
		if colors[equal] < pivot {
			// swap
			colors[smaller], colors[equal] = colors[equal], colors[smaller]
			smaller++
			equal++
		} else if colors[equal] == pivot {
			equal++
		} else {  // colors[equal] > pivot
			colors[equal], colors[larger] = colors[larger], colors[equal]
			larger--
		}
	}
}

type Color int

const (
	Red 	Color = iota
	White
	Blue
)

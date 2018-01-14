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
	i := 0
	j := 0
	n := len(colors) - 1

	pivot := colors[pivotIndex]

	for j <= n {
		if colors[j] < pivot {
			swap(colors, i, j)
			i++
			j++
		} else if colors[j] == pivot {
			j++
		} else { // colors[equal] > pivot
			swap(colors, j, n)
			n--
		}
	}
}

// DutchNationalFlagWithoutPivot sorts the array into three partition
func DutchNationalFlagWithoutPivot(colors []Color) {
	i := 0
	j := 0
	n := len(colors) - 1

	for j < n {
		if colors[j] == Red {
			swap(colors, i, j)
			i++
			j++
		} else if colors[j] == White {
			j++
		} else if colors[j] == Blue {
			swap(colors, j, n)
			n--
		}
	}

}

func swap(colors []Color, a, b int) {
	colors[a], colors[b] = colors[b], colors[a]
}

type Color int

const (
	Red Color = iota
	White
	Blue
)

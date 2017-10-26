package ch14sorting

// MergingTwoSortedArrays merge two sorted array
// It should use the existing array
// One of the array is sufficient to allocate both of them
// Given l1 = { 5, 13, 17, 0, 0, 0, 0, 0 } , m = 3
// and l2 = { 3, 7, 11, 19 }, n = 4
// The result should be 3 5 7 11 13 17 19 0
func MergingTwoSortedArrays(l1 []int, m int, l2 []int, n int) {
	widx := m + n - 1
	midx := m - 1
	nidx := n - 1

	for midx >= 0 && nidx >= 0 {
		if l1[midx] > l2[nidx] {
			l1[widx] = l1[midx]
			midx--
		} else {
			l1[widx] = l2[nidx]
			nidx--
		}
		widx --
	}

	for nidx >= 0 {
		l1[widx] = l2[nidx]
		widx--
		nidx--
	}
}

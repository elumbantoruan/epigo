package ch14sorting

// IntersectionTwoSortedArrays returns an intersection (list) between the two array
// given l1 {2, 3, 3, 5, 5, 6, 7, 7, 8, 12};
// and l2 {5, 5, 6, 8, 8, 9, 10, 10};
// the intersection will be {5,6,8}
func IntersectionTwoSortedArrays(l1, l2 []int) []int {
	l1start := 0
	l1end := len(l1) - 1
	l2start := 0
	l2end := len(l2) - 1

	results := make(map[int]int)

	for l1start <= l1end && l2start <= l2end {
		if l1[l1start] < l2[l2start] {
			l1start++
		} else if l1[l1start] > l2[l2start] {
			l2start++
		} else {
			_, ok := results[l1[l1start]]
			if !ok {
				results[l1[l1start]] = l1[l1start]
			}
			l1start++
			l2start++
		}
	}
	values := []int{}

	for key, _ := range results {
		values = append(values, key)
	}
	return values

}

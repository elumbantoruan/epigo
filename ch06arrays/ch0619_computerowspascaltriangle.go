package ch06arrays

// ComputeRowsPascalTriangle
//					1
//				1		1
//			1		2		1
//		1		3		3		1
//	1		4		6		4		1
func ComputeRowsPascalTriangle(n int) [][]int {
	results := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				val := results[i-1][j-1] + results[i-1][j]
				row = append(row, val)
			}
		}
		results = append(results, row)
	}
	return results
}

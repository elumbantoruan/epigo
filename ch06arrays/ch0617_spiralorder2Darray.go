package ch06arrays

// MatrixInSpiralOrder returns a list of integer
// as a result of spiral movement
// input:
//			c0	c1	c2	c3
//		r0	1	2	3	4
//		r1	5	6	7	8
//		r2	9	10	11	12
//		r3	13	14	15	16
// output:
//		{ 1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10 }
//
func MatrixInSpiralOrder(matrix [][]int) []int {
	spiral := []int{}
	rowStart := 0
	rowEnd := len(matrix) - 1
	colStart := 0
	colEnd := len(matrix[0]) - 1

	for rowStart < rowEnd && colStart < colEnd {
		// colStart: 0 rowStart:0  1,2,3,4
		// colStart: 1 rowStart:1  6,7
		for i := colStart; i <= colEnd; i++ {
			spiral = append(spiral, matrix[rowStart][i])
		}
		// 8, 12, 16
		for i := rowStart + 1; i <= rowEnd; i++ {
			spiral = append(spiral, matrix[i][colEnd])
		}
		// 15, 14
		for i := colEnd - 1; i > colStart; i-- {
			spiral = append(spiral, matrix[rowEnd][i])
		}

		// 13, 9, 5
		for i := rowEnd; i > rowStart; i-- {
			spiral = append(spiral, matrix[i][colStart])
		}
		rowStart++
		rowEnd--
		colStart++
		colEnd--
	}
	return spiral

}

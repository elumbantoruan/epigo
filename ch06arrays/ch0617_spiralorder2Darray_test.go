package ch06arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
func TestMatrixInSpiralOrder(t *testing.T) {
	var matrix [][]int
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	row4 := []int{13, 14, 15, 16}

	matrix = append(matrix, row1)
	matrix = append(matrix, row2)
	matrix = append(matrix, row3)
	matrix = append(matrix, row4)

	spiral := MatrixInSpiralOrder(matrix)
	expected := []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}

	assert.Equal(t, len(expected), len(spiral))
	for i, n := range spiral {
		assert.Equal(t, expected[i], n)
	}
}

package ch06arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate2DArray(t *testing.T) {
	matrix := [][]int{}
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	row4 := []int{13, 14, 15, 16}

	matrix = append(matrix, row1)
	matrix = append(matrix, row2)
	matrix = append(matrix, row3)
	matrix = append(matrix, row4)

	Rotate2DArray(matrix)

	expected := [][]int{}
	expRow1 := []int{13, 9, 5, 1}
	expRow2 := []int{14, 10, 6, 2}
	expRow3 := []int{15, 11, 7, 3}
	expRow4 := []int{16, 12, 8, 4}

	expected = append(expected, expRow1)
	expected = append(expected, expRow2)
	expected = append(expected, expRow3)
	expected = append(expected, expRow4)

	for i := 0; i < len(matrix); i++ {
		mrow := matrix[i]
		erow := expected[i]

		for j := 0; j < len(mrow); j++ {
			assert.Equal(t, mrow[j], erow[j])
		}
	}
}

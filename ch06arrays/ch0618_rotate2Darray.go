package ch06arrays

func Rotate2DArray(matrix [][]int) {
	size := len(matrix) - 1
	for row := 0; row < len(matrix) / 2; row++ {
		for col := row; col < size - row; col++ {
			temp1 := matrix[size - col][row]
			temp2 := matrix[size - row][size - col]
			temp3 := matrix[col][size - row]
			temp4 := matrix[row][col]

			matrix[row][col] = temp1
			matrix[size - col][row] = temp2
			matrix[size - row][size - col] = temp3
			matrix[col][size - row] = temp4
		}
	}
}

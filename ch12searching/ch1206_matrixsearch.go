package ch12searching

// SearchIn2DSortedArray returns a boolean if k is found
// in the matrix
// Given matrix of
//       c0  c1  c2  c3  c4
//   r0	{-1,  2,  4,  4,  6}
//	 r1	{ 1,  5,  5,  9, 21}
//	 r2	{ 3,  6,  6,  9, 22}
//	 r3	{ 3,  6,  8, 10, 24}
//	 r4	{ 6,  8,  9, 12, 25}
//	 r5	{ 8, 10, 12, 13, 40}
// and n: 10
// the return should be true since 10 is found in
// intersection of r3 and c3
func SearchIn2DSortedArray(matrix [][]int, n int) bool {
	rStart := 0
	rEnd := len(matrix) - 1
	cStart := 0
	cEnd := len(matrix[0]) - 1

	for rStart <= rEnd && cStart <= cEnd {
		if matrix[rStart][cEnd] < n {
			rStart++
		} else if matrix[rStart][cEnd] > n {
			cEnd--
		} else if matrix[rStart][cEnd] == n {
			return true
		}
	}
	return false
}

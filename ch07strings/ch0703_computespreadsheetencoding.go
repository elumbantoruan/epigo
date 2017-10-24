package ch07strings

// SpreadsheetEncoding encodes string represented in spreadsheet to a number
// Given columns of A, B, C, AA, AB, ZZ
// it encoded to 1, 2, 3, 27, 28, 702
// There are 26 characters in A .. Z and each can be mapped to an integer
func SpreadsheetEncoding(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = result * 26 + int(s[i] - 'A' + 1)
	}
	return result
}

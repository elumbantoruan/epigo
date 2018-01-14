package ch07strings

//var mapper map[byte]int{
//	'I':1,
//	'V':5,
//	'X':10,
//	'L':50,
//	'C':100,
//	'D':500,
//	'M':1000,
//}

// ConvertRomanToDecimal converts roman to integer value
func ConvertRomanToDecimal(s string) int {
	mapper := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := mapper[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		val := mapper[s[i]]
		if val < mapper[s[i+1]] {
			result -= val
		} else {
			result += val
		}
	}
	return result
}

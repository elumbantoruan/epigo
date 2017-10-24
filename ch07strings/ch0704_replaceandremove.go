package ch07strings

// ReplaceAndRemove replace each 'a' by two 'd'
// and delete each entry containing a 'b
// input of {a, c, d, b, b, c, a}
// output   {d, d, c, d, c, d, d}
func ReplaceAndRemove(runes []rune) {
	writeIdx := 0
	aCount := 0
	// a, c, d, c, a,   c, a
	for i := 0; i < len(runes); i++ {
		if runes[i] != 'b' {
			runes[writeIdx] = runes[i]
			writeIdx++
		}
		if runes[i] == 'a' {
			aCount++
		}
	}

	// backward iteration, replace 'a' with two 'd' from the end
	currentIndex := writeIdx - 1
	writeIdx += aCount -1

	for i := currentIndex; i >= 0; i-- {
		if runes[i] == 'a' {
			runes[writeIdx] = 'd'
			writeIdx --
			runes[writeIdx] = 'd'
			writeIdx --
		} else {
			runes[writeIdx] = runes[i]
			writeIdx--
		}
	}
}

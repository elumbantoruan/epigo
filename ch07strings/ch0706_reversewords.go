package ch07strings

func ReverseWords(runes []rune) {
	reverseChars(0, len(runes)-1, runes)
	start := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			reverseChars(start, i-1, runes)
			start = i + 1
		} else if i == len(runes)-1 {
			reverseChars(start, i, runes)
		}
	}
}

func reverseChars(start, end int, runes []rune) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

package ch13hashtables

// FindTheNearestRepeatedEntries returns the smallest difference between the repeated words
// Given "All", "work", "and", "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no", "results"
// The nearest repeated entries is 2, which is the difference between index 7 and 9 "no", "work", "no"
func FindTheNearestRepeatedEntries(words []string) int {
	entriesIndex := map[string]int{}
	nearestIndex := len(words)
	for i := 0; i < len(words); i++ {
		if v, ok := entriesIndex[words[i]]; ok {
			nearestIndex = min(nearestIndex, i-v)
		}
		entriesIndex[words[i]] = i
	}
	return nearestIndex
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

package ch07strings

// SearchSubstring is to find the index of a substring (s) in an entire text (t)
func SearchSubstring(t, s string) int {
	if len(s) > len(t) {
		return -1
	}

	BASE := 26	// a - z
	tHash := 0	// hash code for t
	sHash := 0	// hash code for s
	powerS := 1	// used for BASE^[s] to roll hash

	for i := 0; i < len(s); i++ {
		if i == 0 {
			powerS = 1
		} else {
			powerS = powerS * BASE
		}
		sHash = sHash * BASE + int(s[i])
		tHash = tHash * BASE + int(t[i])
	}

	// this is to check if the end of the last character of substring is not located the end
	// of entire text
	// t: "this is a substring test"
	// s: "substring"
	for i := len(s); i < len(t); i++ {
		// check if the two substrings are actually equal or not, to protect
		// against hash collision
		if tHash == sHash && t[i - len(s):i] == s {
			return i - len(s)
		}

		// uses rolling hash to compute new hash code
		tHash = tHash - int(t[i - len(s)]) * powerS
		tHash = tHash * BASE + int(t[i])
	}

	// this check is used if the end of the last character of a substring is located at the end of t
	// t: "this is a test for substring"
	// s: "substring"
	if tHash == sHash && t[len(t) - len(s):] == s {
		return len(t) - len(s)
	}

	return -1

}

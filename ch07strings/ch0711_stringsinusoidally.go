package ch07strings

import "strings"

func StringSinusoidally(s string) string {
	top := []byte{}
	for i := 1; i < len(s); i += 4 {
		top = append(top, s[i])
	}

	middle := []byte{}
	for i := 0; i < len(s); i += 2 {
		middle = append(middle, s[i])
	}

	bottom := []byte{}
	for i := 3; i < len(s); i += 4 {
		bottom = append(bottom, s[i])
	}

	ss := []string{
		string(top),
		string(middle),
		string(bottom),
	}

	return strings.Join(ss, "")
}

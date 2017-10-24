package ch07strings

import (
	"strconv"
)

func GetAllValidIPAddress(s string) []string {
	if len(s) == 0 || len(s) < 4 || len(s) > 12 {
		panic("not a valid length")
	}
	results := []string{}

	for i := 1; i < len(s) && i < 4; i++ {
		first := s[0:i]
		if isValidPart(first) {
			for j := 1; j < len(s) && j < 4; j++ {
				second := s[i: i + j]
				if isValidPart(second) {
					for k := 1; k < len(s) && k < 4; k++ {
						if i + j + k < len(s) {
							third := s[i + j: i + j + k]
							forth := s[i + j + k:]
							if isValidPart(third) && isValidPart(forth) {
								results = append(results, first + "." + second + "." + third + "." + forth)
							}
						}
					}
				}
			}
		}
	}
	return results
}

func isValidPart(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	v, _ := strconv.Atoi(s)
	if v > 255 {
		return false
	}
 	return true

}

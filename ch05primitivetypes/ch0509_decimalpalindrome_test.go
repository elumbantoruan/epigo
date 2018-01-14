package ch05primitivetypes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	n := 232
	palindrome := IsPalindrome(n)
	assert.Equal(t, true, palindrome)
}

func ExampleIsPalindrome() {
	n := 232
	palindrome := IsPalindrome(n)
	fmt.Println(palindrome)
	// Output: true

}

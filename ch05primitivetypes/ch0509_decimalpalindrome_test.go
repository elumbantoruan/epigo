package ch05primitivetypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
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
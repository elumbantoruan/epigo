package ch05primitivetypes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseDigit(t *testing.T) {
	n := 23
	reversed := ReverseDigit(n)
	assert.Equal(t, 32, reversed)
}

func ExampleReverseDigit() {
	n := 23
	reversed := ReverseDigit(n)
	fmt.Println(reversed)
	// Output: 32
}

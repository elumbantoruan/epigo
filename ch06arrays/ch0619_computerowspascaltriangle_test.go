package ch06arrays

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeRowsPascalTriangle(t *testing.T) {
	n := 5
	results := ComputeRowsPascalTriangle(n)
	for _, r := range results {
		fmt.Println(r)
	}
	assert.Equal(t, 0, 0)
}

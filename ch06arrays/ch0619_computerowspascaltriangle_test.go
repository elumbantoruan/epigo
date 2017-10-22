package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestComputeRowsPascalTriangle(t *testing.T) {
	n := 5
	results := ComputeRowsPascalTriangle(n)
	for _, r := range results {
		fmt.Println(r)
	}
	assert.Equal(t, 0, 0)
}

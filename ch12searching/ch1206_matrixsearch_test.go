package ch12searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
{-1,  2,  4,  4,  6}
{ 1,  5,  5,  9, 21}
{ 3,  6,  6,  9, 22}
{ 3,  6,  8, 10, 24}
{ 6,  8,  9, 12, 25}
{ 8, 10, 12, 13, 40}
*/

func TestSearchIn2DSortedArray(t *testing.T) {
	matrix := [][]int{
		{-1, 2, 4, 4, 6},
		{1, 5, 5, 9, 21},
		{3, 6, 6, 9, 22},
		{3, 6, 8, 10, 24},
		{6, 8, 9, 12, 25},
		{8, 10, 12, 13, 40},
	}
	n := 10
	b := SearchIn2DSortedArray(matrix, n)
	e := true
	assert.Equal(t, e, b)

	n = 100
	b = SearchIn2DSortedArray(matrix, n)
	e = false
	assert.Equal(t, e, b)
}

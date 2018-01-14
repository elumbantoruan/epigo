package ch06arrays

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDutchNationalFlag(t *testing.T) {
	colors := []Color{
		Blue,
		Red,
		Red,
		Blue,
		White,
		Red,
		White,
	}

	pvtIdx := 1
	DutchNationalFlag(colors, pvtIdx)

	assert.Equal(t, colors[0], Red)
	assert.Equal(t, colors[1], Red)
	assert.Equal(t, colors[2], Red)

	for _, c := range colors {
		fmt.Printf("%v\n", c)
	}
}

func TestDutchNationalFlagWithoutPivot(t *testing.T) {
	colors := []Color{
		Blue,
		Red,
		Red,
		Blue,
		White,
		Red,
		White,
	}

	DutchNationalFlagWithoutPivot(colors)

	expected := []Color{
		Red,
		Red,
		Red,
		White,
		White,
		Blue,
		Blue,
	}

	for i, e := range colors {
		assert.Equal(t, expected[i], e)
	}
}

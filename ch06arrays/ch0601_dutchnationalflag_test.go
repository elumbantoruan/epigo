package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestDutchNationalFlag(t *testing.T) {
	colors := []Color {
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

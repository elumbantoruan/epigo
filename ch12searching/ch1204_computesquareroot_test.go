package ch12searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeSquareRoot(t *testing.T) {
	n := 16
	r := ComputeSquareRoot(n)
	e := 4
	assert.Equal(t, e, r)

	n = 121
	r = ComputeSquareRoot(n)
	e = 11
	assert.Equal(t, e, r)

}

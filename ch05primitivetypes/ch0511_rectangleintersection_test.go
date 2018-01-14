package ch05primitivetypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangleIntersection(t *testing.T) {
	r1 := Rectangular{X: 20, Y: 10, Width: 50, Height: 30}
	r2 := Rectangular{X: 40, Y: 5, Width: 10, Height: 45}
	ir := RectangleIntersection(r1, r2)
	t.Log(ir.String())
	assert.Equal(t, float64(40), ir.X)
	assert.Equal(t, float64(10), ir.Y)
	assert.Equal(t, float64(10), ir.Width)
	assert.Equal(t, float64(30), ir.Height)
}

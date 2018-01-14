package ch05primitivetypes

import (
	"fmt"
	"math"
)

// RectangleIntersection computes the intersection between the two rectangulars
func RectangleIntersection(r1, r2 Rectangular) *Rectangular {
	if !isIntersect(r1, r2) {
		return nil
	}
	x := math.Max(r1.X, r2.X)
	y := math.Max(r1.Y, r2.Y)
	w := math.Min(r1.X+r1.Width, r2.X+r2.Width) - math.Max(r1.X, r2.X)
	h := math.Min(r1.Y+r1.Height, r2.Y+r2.Height) - math.Max(r1.Y, r2.Y)

	return &Rectangular{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

func isIntersect(r1, r2 Rectangular) bool {
	return r1.X <= r2.X+r2.Width && r1.X+r1.Width >= r2.X &&
		r1.Y <= r2.Y+r2.Height && r1.Y+r1.Height >= r2.Y
}

// Rectangular represents a rectangular structure with x,y,width, and height
type Rectangular struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// String returns values of rectangular in string literals
func (r *Rectangular) String() string {
	if r == nil {
		return "nil"
	}
	return fmt.Sprintf("X: %v Y: %v Width: %v Height %v",
		r.X, r.Y, r.Width, r.Height)
}

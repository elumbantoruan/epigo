package ch14sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergingInterval(t *testing.T) {
	existingIntervals := []Interval{
		{-4, -1},
		{0, 2},
		{3, 6},
		{7, 9},
		{11, 12},
		{14, 17},
	}
	newInterval := Interval{1, 8}

	results := MergingInterval(existingIntervals, newInterval)
	//assert.True(t, results != nil)
	expected := []*Interval{
		{-4, -1},
		{0, 9},
		{11, 12},
		{14, 17},
	}
	for i, m := range results {
		assert.Equal(t, expected[i], m)
	}
}

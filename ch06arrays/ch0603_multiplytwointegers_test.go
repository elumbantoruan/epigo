package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiply_25times25(t *testing.T) {
	n1 := []int {2,5}
	n2 := []int {2,5}
	result := Multiply(n1, n2)
	expected := []int {6,2,5}
	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}

func TestMultiply_25timesminus25(t *testing.T) {
	n1 := []int {2,5}
	n2 := []int {-2,5}
	result := Multiply(n1, n2)
	expected := []int {-6,2,5}
	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}

func TestMultiply_2times2(t *testing.T) {
	n1 := []int {2}
	n2 := []int {2}
	result := Multiply(n1, n2)
	expected := []int {4}
	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}

func TestMultiply_2times9(t *testing.T) {
	n1 := []int {2}
	n2 := []int {9}
	result := Multiply(n1, n2)
	expected := []int {1,8}
	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}

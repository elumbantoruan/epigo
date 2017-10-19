package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEvenOddMerge(t *testing.T) {
	l := FromArray([]interface{} {0, 1, 2, 3, 4, 5})
	r := EvenOddMerge(l)
	expected := []interface{} {0, 2, 4, 1, 3, 5}
	actual := r.ToArray()

	for i, e := range actual {
		assert.Equal(t, expected[i], e)
	}
}

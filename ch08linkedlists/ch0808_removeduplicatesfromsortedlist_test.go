package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	l := FromArray([]interface{} { 2, 2, 3, 5, 7, 11, 11 })
	l.Print()

	r := RemoveDuplicatesFromSortedList(l)

	actual := r.ToArray()
	expected := []interface{} { 2, 3, 5, 7, 11}

	for i, e := range actual {
		assert.Equal(t, expected[i], e)
	}
}

package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddListBasedIntegers(t *testing.T) {
	l1 := FromArray([]interface{} {3,1,4})
	l2 := FromArray([]interface{} {7,0,9})

	r := AddListBasedIntegers(l1, l2)

	actual := r.ToArray()
	expected := []interface{} {0,2,3,1}

	for i, e := range actual {
		assert.Equal(t, expected[i], e)
	}
}

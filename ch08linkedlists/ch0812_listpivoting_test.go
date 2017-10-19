package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListPivoting(t *testing.T) {
	l := FromArray([]interface{} { 3, 2, 2, 11, 7, 5, 11 })
	k := 7
	r := ListPivoting(l, k)

	expected := []interface{} { 3, 2, 2, 5, 7, 11, 11 }
	actual := r.ToArray()

	for i, e := range actual {
		assert.Equal(t, expected[i], e)
	}
}

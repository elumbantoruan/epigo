package ch08linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKthLastElement(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	l := FromArray(arr)
	l.Print()

	k := 4
	r := RemoveKthLastElement(l, k)

	actual := r.ToArray()
	expected := []interface{}{1, 2, 3, 4, 6, 7, 8}

	assert.Equal(t, expected, actual)

}

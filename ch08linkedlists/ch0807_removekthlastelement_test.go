package ch08linkedlists

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveKthLastElement(t *testing.T) {
	arr := []interface{} {1,2,3,4,5,6,7,8}
	l := FromArray(arr)
	l.Print()

	k := 4
	r := RemoveKthLastElement(l, k)

	r.Print()
	assert.True(t, r != nil, "result should not be nil")

}

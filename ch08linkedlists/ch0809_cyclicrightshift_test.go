package ch08linkedlists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyclicRightShift(t *testing.T) {
	l := FromArray([]interface{}{2, 3, 5, 3, 2})
	k := 3
	fmt.Println("\nCyclicRightShift input:")
	l.Print()
	r := CyclicRightShift(l, k)
	fmt.Println("\nCyclicRightShift output:")
	r.Print()

	actual := r.ToArray()
	expected := []interface{}{5, 3, 2, 2, 3}

	for i, e := range actual {
		assert.Equal(t, expected[i], e)
	}

}

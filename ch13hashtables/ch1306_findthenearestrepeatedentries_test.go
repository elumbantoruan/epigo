package ch13hashtables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheNearestRepeatedEntries(t *testing.T) {
	words := []string{"All", "work", "and", "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no", "results"}
	result := FindTheNearestRepeatedEntries(words)
	expected := 2
	assert.Equal(t, expected, result)
}

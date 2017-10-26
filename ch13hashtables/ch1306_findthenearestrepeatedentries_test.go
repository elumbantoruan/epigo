package ch13hashtables

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindTheNearestRepeatedEntries(t *testing.T) {
	words := []string{"All", "work", "and", "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no", "results"}
	result := FindTheNearestRepeatedEntries(words)
	expected := 2
	assert.Equal(t, expected, result)
}

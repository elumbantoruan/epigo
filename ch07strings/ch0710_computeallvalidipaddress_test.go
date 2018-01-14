package ch07strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllValidIPAddress(t *testing.T) {
	s := "19216811"
	results := GetAllValidIPAddress(s)
	assert.Contains(t, results, "192.168.1.1")
}

package ch07strings

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetAllValidIPAddress(t *testing.T) {
	s := "19216811"
	results := GetAllValidIPAddress(s)
	assert.Contains(t, results, "192.168.1.1")
}

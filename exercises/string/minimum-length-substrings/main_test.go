package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinLengthSubstring(t *testing.T) {
	assert.Equal(t, -1, minLengthSubstring("zzzz", "a"))
	assert.Equal(t, -1, minLengthSubstring("", "fd"))
	assert.Equal(t, -1, minLengthSubstring("dcbefebce", ""))
	assert.Equal(t, -1, minLengthSubstring("bfbeadbcbcbfeaaeefcddcccbbbfaaafdbebedddf", "cbccfafebccdccebdd"))
	assert.Equal(t, 2, minLengthSubstring("dcffd", "fd"))
	assert.Equal(t, 4, minLengthSubstring("ADOBECODEBANC", "ABC"))
	assert.Equal(t, 5, minLengthSubstring("dcbefebce", "fd"))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOperations(t *testing.T) {
	assert.Equal(t, 0, minOperations([]int{1}))
	assert.Equal(t, 2, minOperations([]int{3, 1, 2}))
	assert.Equal(t, 1, minOperations([]int{1, 2, 5, 4, 3}))
}

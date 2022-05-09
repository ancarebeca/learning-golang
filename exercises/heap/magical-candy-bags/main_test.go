package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxCandies(t *testing.T) {
	assert.Equal(t, 14, maxCandies([]int{2, 1, 7, 4, 2}, 3))
	assert.Equal(t, 228, maxCandies([]int{19, 78, 76, 72, 48, 8, 24, 74, 29}, 3))
}

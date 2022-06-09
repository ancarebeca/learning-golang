package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	assert.Equal(t, []int{0, 2}, searchRange([]int{3, 3, 3}, 3))

	assert.Equal(t, []int{0, 0}, searchRange([]int{1}, 1))

	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	assert.Equal(t, []int{0, 1}, searchRange([]int{2, 2}, 2))

	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 3))

	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

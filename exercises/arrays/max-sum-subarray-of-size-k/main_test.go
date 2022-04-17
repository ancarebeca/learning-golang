package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxSum(t *testing.T) {
	assert.Equal(t, 8, getMaxSum([]int{2, 5, 2, 1, 7}, 2))
	assert.Equal(t, 10, getMaxSum([]int{9, 1, 2, 1, 7}, 2))
	assert.Equal(t, 10, getMaxSum([]int{3, 1, 8, 2, 7}, 2))
	assert.Equal(t, 18, getMaxSum([]int{3, 1, 8, 2, 7}, 4))
	assert.Equal(t, 13, getMaxSum([]int{5, 1, 5, 2, 1}, 4))
	assert.Equal(t, 14, getMaxSum([]int{3, 1, 5, 2, 7, 1, 4}, 3))
}

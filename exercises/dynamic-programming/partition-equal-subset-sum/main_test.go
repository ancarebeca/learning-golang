package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanPartitionBruteForce(t *testing.T) {
	assert.False(t, canPartitionBruteForce([]int{2, 6, 7}))
	assert.True(t, canPartitionBruteForce([]int{1, 2, 3, 4}))
	assert.True(t, canPartitionBruteForce([]int{1, 1, 3, 4, 7}))
	assert.True(t, canPartitionBruteForce([]int{1, 5, 11, 5}))

}

func TestCanPartitionTopDown(t *testing.T) {
	assert.True(t, canPartitionTopDown([]int{1, 6, 11, 6}))
	assert.False(t, canPartitionTopDown([]int{2, 6, 7}))
	assert.True(t, canPartitionTopDown([]int{1, 1, 3, 4, 7}))
}

func TestCanPartitionBottomUp(t *testing.T) {
	assert.True(t, canPartitionBottomUp([]int{1, 6, 11, 6}))
	assert.False(t, canPartitionBottomUp([]int{2, 6, 7}))
	assert.True(t, canPartitionBottomUp([]int{1, 1, 3, 4, 7}))
}

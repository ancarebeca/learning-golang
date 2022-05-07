package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	assert.Equal(t, 0, coinChangeBruteForce([]int{1}, 0))
	assert.Equal(t, -1, coinChangeBruteForce([]int{2}, 3))
	assert.Equal(t, 2, coinChangeBruteForce([]int{1, 2, 5}, 6))
}

func TestCoinChangeMemoization(t *testing.T) {
	assert.Equal(t, 0, coinChangeMemoization([]int{1}, 0))
	assert.Equal(t, -1, coinChangeMemoization([]int{2}, 3))
	assert.Equal(t, 2, coinChangeMemoization([]int{1, 2, 5}, 6))
	assert.Equal(t, -1, coinChangeMemoization([]int{4, 5}, 11))
	assert.Equal(t, 20, coinChangeMemoization([]int{1, 2, 5}, 100))
	assert.Equal(t, 20, coinChangeMemoization([]int{186, 419, 83, 408}, 6249))
	assert.Equal(t, 3, coinChangeMemoization([]int{1, 5, 10, 21, 25}, 63))
	assert.Equal(t, 2, coinChangeMemoization([]int{2, 3, 4}, 6))

}

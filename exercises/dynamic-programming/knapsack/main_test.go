package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveKnapsackBottomUp(t *testing.T) {
	assert.Equal(t, 17, solveKnapsackBottomUp([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))
	assert.Equal(t, 22, solveKnapsackBottomUp([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7))
	assert.Equal(t, 0, solveKnapsackBottomUp([]int{1, 6, 10, 16}, []int{2, 2, 3, 5}, 1))
	assert.Equal(t, 5, solveKnapsackBottomUp([]int{3, 1, 2}, []int{1, 1, 2}, 3))

}

func TestSolveKnapsackTopDown(t *testing.T) {
	assert.Equal(t, 17, solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))
	assert.Equal(t, 22, solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7))
	assert.Equal(t, 0, solveKnapsack([]int{1, 6, 10, 16}, []int{2, 2, 3, 5}, 1))
	assert.Equal(t, 5, solveKnapsack([]int{3, 1, 2}, []int{1, 1, 2}, 3))

}

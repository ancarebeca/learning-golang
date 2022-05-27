package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, 43, minPathSum([][]int{
		{9, 4, 9, 9},
		{6, 7, 6, 4},
		{8, 3, 3, 7},
		{7, 4, 9, 10},
	}))

	assert.Equal(t, 12, minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))

	assert.Equal(t, 22, minPathSum([][]int{
		{9, 1, 4, 8},
	}))

	assert.Equal(t, 5, minPathSum([][]int{
		{1},
		{4},
	}))
}

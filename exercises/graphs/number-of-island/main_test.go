package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIslands(t *testing.T) {
	assert.Equal(t, 2, numIslands([][]byte{
		{1, 1, 0},
		{0, 0, 1},
	}))

	assert.Equal(t, 2, numIslands([][]byte{

		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}))

	assert.Equal(t, 1, numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}

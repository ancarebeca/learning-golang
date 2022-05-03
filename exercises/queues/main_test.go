package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPositions(t *testing.T) {
	assert.Equal(t, []int{5, 6, 4, 1, 2}, findPositions([]int{1, 2, 2, 3, 4, 5}, 5))
	assert.Equal(t, []int{2, 5, 10, 13}, findPositions([]int{2, 4, 2, 4, 3, 1, 2, 2, 3, 4, 3, 4, 4}, 4))
}

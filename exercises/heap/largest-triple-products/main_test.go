package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxProduct(t *testing.T) {
	assert.Equal(t, []int{-1, -1, 6, 24, 60}, findMaxProduct([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{-1, -1, 4, 4, 8}, findMaxProduct([]int{2, 1, 2, 1, 2}))
}

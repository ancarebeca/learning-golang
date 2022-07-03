package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptimize(t *testing.T) {
	assert.Equal(t, []int{6, 7}, optimize([]int{3, 7, 5, 6, 2}))
	assert.Equal(t, []int{5, 5, 5, 5}, optimize([]int{5, 5, 5, 5, 5, 5, 5}))

}

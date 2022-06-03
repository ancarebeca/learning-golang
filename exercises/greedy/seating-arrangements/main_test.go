package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOverallAwkwardness(t *testing.T) {
	assert.Equal(t, 2, minOverallAwkwardness([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 2, minOverallAwkwardness([]int{1, 2, 3}))
	assert.Equal(t, 3, minOverallAwkwardness([]int{3, 3, 3, 6}))
	assert.Equal(t, 0, minOverallAwkwardness([]int{3, 3, 3}))
	assert.Equal(t, 1, minOverallAwkwardness([]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6}))
}

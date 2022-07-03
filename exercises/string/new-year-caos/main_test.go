package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	assert.Equal(t, "4", MinimumBribes([]int32{1, 2, 5, 3, 4, 7, 8, 6}))
	assert.Equal(t, "3", MinimumBribes([]int32{2, 1, 5, 3, 4}))
	assert.Equal(t, "Too chaotic", MinimumBribes([]int32{2, 5, 1, 3, 4}))
	assert.Equal(t, "Too chaotic", MinimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4}))
	assert.Equal(t, "7", MinimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4}))
	assert.Equal(t, "4", MinimumBribes([]int32{1, 2, 5, 3, 4, 7, 8, 6}))

}

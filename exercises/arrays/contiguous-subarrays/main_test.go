package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubarrays(t *testing.T) {
	assert.Equal(t, []int{1}, countSubarrays([]int{1}))

	assert.Equal(t, []int{1, 2}, countSubarrays([]int{1, 3}))

	assert.Equal(t, []int{2, 1}, countSubarrays([]int{3, 1}))

	assert.Equal(t, []int{1, 3, 1}, countSubarrays([]int{3, 4, 1}))

	assert.Equal(t, []int{2, 1, 3}, countSubarrays([]int{6, 4, 7}))

	assert.Equal(t, []int{3, 3, 3}, countSubarrays([]int{6, 6, 6}))

	assert.Equal(t, []int{1, 3, 1, 5, 1}, countSubarrays([]int{3, 4, 1, 6, 2}))

	assert.Equal(t, []int{1, 4, 1, 2}, countSubarrays([]int{3, 7, 1, 6}))
}

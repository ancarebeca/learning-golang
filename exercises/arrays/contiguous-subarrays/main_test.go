package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreTheyEqual(t *testing.T) {
	assert.Equal(t, []int{1}, countSubarrays([]int{1}))

	assert.Equal(t, []int{1, 2}, countSubarrays([]int{1, 3}))

	assert.Equal(t, []int{2, 1}, countSubarrays([]int{3, 1}))

	assert.Equal(t, []int{1, 3, 1}, countSubarrays([]int{3, 4, 1}))

	assert.Equal(t, []int{1, 3, 1, 5, 1}, countSubarrays([]int{3, 4, 1, 6, 2}))
}

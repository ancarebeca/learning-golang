package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreTheyEqual(t *testing.T) {
	assert.Equal(t, []int{1}, countSubarrays([]int{1}))

	assert.Equal(t, []int{1, 3}, countSubarrays([]int{1, 2}))
}

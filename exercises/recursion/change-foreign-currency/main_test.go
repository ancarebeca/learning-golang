package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanGetExactChangeBruteForce(t *testing.T) {
	assert.True(t, canGetExactChangeBruteForce(75, []int{4, 17, 29}))
	assert.False(t, canGetExactChangeBruteForce(94, []int{5, 10, 25, 100, 200}))
	assert.True(t, canGetExactChangeBruteForce(7, []int{1, 3, 4, 5}))
	assert.True(t, canGetExactChangeBruteForce(9, []int{5, 3, 2}))
}

func TestCanGetExactChangeBruteMemoization(t *testing.T) {
	assert.True(t, canGetExactChangeVersion2(75, []int{4, 17, 29}))
	assert.False(t, canGetExactChangeVersion2(94, []int{5, 10, 25, 100, 200}))
	assert.True(t, canGetExactChangeVersion2(7, []int{1, 3, 4, 5}))
}

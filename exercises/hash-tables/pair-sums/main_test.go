package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairSums(t *testing.T) {

	assert.Equal(t, int32(0), pairSums(6, 5, []int32{}))
	assert.Equal(t, int32(0), pairSums(6, 5, []int32{6}))
	assert.Equal(t, int32(0), pairSums(6, 5, []int32{1, 1}))
	assert.Equal(t, int32(0), pairSums(6, 5, []int32{1, 1}))
	assert.Equal(t, int32(0), pairSums(6, 5, []int32{10, 10}))

	assert.Equal(t, int32(2), pairSums(6, 5, []int32{1, 2, 3, 4, 4}))
	assert.Equal(t, int32(4), pairSums(6, 5, []int32{1, 5, 3, 3, 3}))
	assert.Equal(t, int32(2), pairSums(6, 5, []int32{1, 2, 3, 4, 3}))
}

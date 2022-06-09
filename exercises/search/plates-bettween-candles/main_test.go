package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlatesBetweenCandles(t *testing.T) {
	assert.Equal(t, []int{2, 3}, platesBetweenCandles("**|**|***|*", [][]int{{2, 5}, {5, 9}}))
	assert.Equal(t, []int{2}, platesBetweenCandles("**|**|***|*", [][]int{{1, 7}}))
	assert.Equal(t, []int{9, 0, 0, 0, 0}, platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
}

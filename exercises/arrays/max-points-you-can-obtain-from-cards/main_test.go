package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScoreUsingSlidingWindow(t *testing.T) {
	assert.Equal(t, 12, maxScoreUsingSlidingWindow([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	//assert.Equal(t, 4, maxScoreUsingSlidingWindow([]int{2, 2, 2}, 2))
	//assert.Equal(t, 55, maxScoreUsingSlidingWindow([]int{9, 7, 7, 9, 7, 7, 9}, 7))

}

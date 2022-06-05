package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOverallAwkwardness(t *testing.T) {
	assert.Equal(t, 2, maxNumberOfFamilies(2, [][]int{
		{1, 2},
		{1, 3},
		{1, 8},
		{2, 6},
		{3, 1},
		{3, 10},
	}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	assert.Equal(t, 2, findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	assert.Equal(t, 3, findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	assert.Equal(t, 1, findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	assert.Equal(t, 4, orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

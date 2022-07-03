package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterItemsInContainers(t *testing.T) {
	assert.Equal(t, []int{2, 3}, counter("|**|*|*", []int{1, 1}, []int{5, 6}))
	assert.Equal(t, []int{2, 1}, counter("*|*|*|", []int{1, 4}, []int{6, 6}))
}

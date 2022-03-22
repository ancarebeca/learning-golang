package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestAreTheyEqual(t *testing.T) {
	assert.True(t, areTheyEqual([]int{1, 3, 2, 4}  , []int{1, 2, 3, 4}))
}

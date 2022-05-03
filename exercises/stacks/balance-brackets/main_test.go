package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	assert.True(t, isBalanced("{[()]}"))
	assert.True(t, isBalanced("{}()"))
	assert.False(t, isBalanced("{(})"))
	assert.False(t, isBalanced(")"))
	assert.False(t, isBalanced("{([)}()"))
	assert.False(t, isBalanced("{(["))
	assert.False(t, isBalanced("("))
}

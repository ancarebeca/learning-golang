package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {

	assert.Equal(t, "xbacbca", encrypt("abcxcba"))
	assert.Equal(t, "bac", encrypt("abc"))
	assert.Equal(t, "bacd", encrypt("abcd"))
	assert.Equal(t, "a", encrypt("a"))
	assert.Equal(t, "eafcobok", encrypt("facebook"))

}

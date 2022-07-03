package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuperDigit(t *testing.T) {
	assert.Equal(t, int32(3), SuperDigit("148", 3))
	assert.Equal(t, int32(9), SuperDigit("123", 3))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortString(t *testing.T) {
	assert.Equal(t, "YES", gridChallenge([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}))
	assert.Equal(t, "NO", gridChallenge([]string{"mpxz", "abcd", "wlmf"}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotationalCipher(t *testing.T) {
	assert.Equal(t, "nopqrstuvwxyzABCDEFGHIJKLM9012345678", rotationalCipher("abcdefghijklmNOPQRSTUVWXYZ0123456789", 39))
	assert.Equal(t, "okffng-Qwvb", rotationalCipher("middle-Outz", 2))
	assert.Equal(t, "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj", rotationalCipher("Always-Look-on-the-Bright-Side-of-Life", 5))
}

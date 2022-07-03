package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruckTour(t *testing.T) {
	// petrol 0  distance 1
	assert.Equal(t, int32(1), TruckTour([][]int32{{1, 5}, {10, 3}, {3, 4}}))

	assert.Equal(t, int32(0), TruckTour([][]int32{{2, 1}, {1, 1}, {4, 3}, {3, 2}, {2, 1}}))

}

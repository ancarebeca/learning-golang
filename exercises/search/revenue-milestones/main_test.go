package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMilestoneDays(t *testing.T) {
	assert.Equal(t, []int{-1}, getMilestoneDays([]int{50}, []int{100}))
	assert.Equal(t, []int{1, 2}, getMilestoneDays([]int{100, 200}, []int{100, 200}))
	assert.Equal(t, []int{1, 2}, getMilestoneDays([]int{100, 200}, []int{100, 200}))
	assert.Equal(t, []int{4, 2, 4}, getMilestoneDays([]int{100, 300, 200, 400, 500}, []int{700, 300, 900}))
	assert.Equal(t, []int{4, 6, 10}, getMilestoneDays([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{100, 200, 500}))
	assert.Equal(t, []int{2, 4, 4, 5}, getMilestoneDays([]int{100, 200, 300, 400, 500}, []int{300, 800, 1000, 1400}))
	assert.Equal(t, []int{5, 4, 2, 3, 2}, getMilestoneDays([]int{700, 800, 600, 400, 600, 700}, []int{3100, 2200, 800, 2100, 1000}))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxProduct(t *testing.T) {
	h := maxHeap{[]int{100, 19, 36, 17, 3, 25, 1, 2, 7}}
	h.push(200)
	assert.Equal(t, []int{200, 100, 36, 17, 19, 25, 1, 2, 7, 3}, h.elements)

	h.push(4)
	assert.Equal(t, []int{200, 100, 36, 17, 19, 25, 1, 2, 7, 3, 4}, h.elements)

	h.pop()
	assert.Equal(t, []int{100, 19, 36, 17, 4, 25, 1, 2, 7, 3}, h.elements)

	h2 := maxHeap{[]int{100, 19, 36}}
	h2.pop()
	h2.pop()
	assert.Equal(t, []int{19}, h2.elements)

	h2.pop()
	assert.Equal(t, []int{}, h2.elements)

	h3 := maxHeap{[]int{}}
	h3.push(200)
	h3.push(19)
	h3.push(36)
	assert.Equal(t, []int{200, 19, 36}, h3.elements)
}

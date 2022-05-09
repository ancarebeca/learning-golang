package main

import (
	"fmt"
	"math"
)

type maxHeap struct {
	elements []int
}

func (h maxHeap) leftChild(parent int) int {
	return (parent * 2) + 1
}

func (h maxHeap) rightChild(parent int) int {
	return (parent * 2) + 2
}

func (h maxHeap) parent(child int) int {
	return (child - 1) / 2
}

func (h *maxHeap) push(item int) {
	h.elements = append(h.elements, item)

	i := len(h.elements) - 1
	for h.elements[h.parent(i)] < h.elements[i] {
		h.swap(h.parent(i), i)
		i = h.parent(i)
	}
}

func (h *maxHeap) pop() int {
	// take the top of the heap
	ele := h.elements[0]

	// take the last element and make it the top of the heap
	h.elements[0] = h.elements[len(h.elements)-1]

	// rearrange the heap, given the index to start with
	h.rearrange(0)

	return ele
}

func (h maxHeap) swap(i, j int) {
	temp := h.elements[i]
	h.elements[i] = h.elements[j]
	h.elements[j] = temp
}

func (h maxHeap) rearrange(i int) {
	// assume the given index is the largest
	largest := i

	// get the left and right child indices, and the size of the heap
	left, right, size := h.leftChild(i), h.rightChild(i), len(h.elements)

	// if left child is larger than the current largest,
	// set the left to be the largest
	if left < size && h.elements[left] > h.elements[largest] {
		largest = left
	}
	// if right child is larger than the current largest,
	// set the right to be the largest
	if right < size && h.elements[right] > h.elements[largest] {
		largest = right
	}

	// if the largest has changed, swap the positions and recurse
	if largest != i {
		h.swap(i, largest)
		h.rearrange(largest)
	}
}

func maxCandies(arr []int, k int) int {
	h := maxHeap{elements: []int{}}
	for _, element := range arr {
		h.push(element)
	}
	output := 0
	for i := 0; i < k; i++ {
		value := h.pop()
		output += value
		refillWith := int(math.Floor(float64(value) / 2))
		h.push(refillWith)
	}
	return output
}

func main() {
	fmt.Println(maxCandies([]int{2, 1, 7, 4, 2}, 3))
}

package main

type maxHeap struct {
	elements []int
}

func (h maxHeap) leftChild(parent int) int {
	return (2 * parent) + 1
}

func (h maxHeap) rightChild(parent int) int {
	return (2 * parent) + 2
}

func (h maxHeap) parent(child int) int {
	return (child - 1) / 2
}

func (h *maxHeap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *maxHeap) push(item int) {
	// append the new element to the array
	h.elements = append(h.elements, item)

	// starting from the index of the last element
	// swap the index and the index its parent if it's larger
	i := len(h.elements) - 1
	for h.elements[i] > h.elements[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *maxHeap) pop() (ele int) {
	// take the top of the heap
	ele = h.elements[0]
	// take the last element and make it the top of the heap
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	// rearrange the heap, given the index to start with
	h.rearrange(0)
	return
}

func (h *maxHeap) rearrange(i int) {
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

func main() {

}

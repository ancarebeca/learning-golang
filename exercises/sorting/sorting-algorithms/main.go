package main

func bubbleSort(arr []int) []int {
	// Traverse through all array elements
	for i := 0; i < len(arr); i++ {
		swapped := false
		// Last i elements are already in place
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				continue
			}

			// traverse the array from 0 to n-i-1
			// swap if the element found is greater
			// than the next element
			temp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = temp
			swapped = true
		}

		// If no two elements were swapped
		// by inner loop, then break
		if swapped == false {
			break
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	// Traverse through 1 to len(arr)
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		// Move elements of arr[0..i-1], that are
		// greater than key, to one position ahead
		// of their current position
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2

	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var l, r int
	merged := []int{}

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			merged = append(merged, left[l])
			l++
		} else {
			merged = append(merged, right[r])
			r++
		}
	}

	for ; l < len(left); l++ {
		merged = append(merged, left[l])
	}
	for ; r < len(right); r++ {
		merged = append(merged, right[r])
	}
	return merged
}

func quickSort(arr []int) []int {
	return quickSortRec(arr, 0, len(arr)-1)
}

func quickSortRec(arr []int, lb, ub int) []int {
	if lb < ub {
		// Find the pivot index of a lower bound of an array
		var pivot = partition(arr, lb, ub)

		// Apply Divide and conquer strategy
		// to sort the left side and right side of an array
		// respective to the pivot position

		// Left hand side array
		quickSortRec(arr, lb, pivot-1)

		// Right hand side array
		quickSortRec(arr, pivot+1, ub)
	}
	return arr
}

func partition(arr []int, lb, ub int) int {
	// Pick a lowest bound element as an pivot value
	pivot := arr[lb]
	start := lb
	end := ub

	for start < end {

		// Increment the index value of "start"
		// till the left most values should be less than or equal to the pivot value
		for arr[start] <= pivot && start < ub {
			start++
		}

		// Decrement the index value of "end"
		// till the right most values should be greater than to the pivot value
		for arr[end] > pivot && end > lb {
			end--
		}

		// Interchange the values of present
		// in the index start and end of an array only if start is less than end
		if start < end {
			arr = swap(arr, start, end)
		}
	}

	// Interchange the element in end's poisition to the lower bound of an array
	// and place the Pivot element to the end's position
	arr = swap(arr, lb, end)

	// Finally return the appropriate index position of the pivot element
	return end
}

func swap(arr []int, start int, end int) []int {
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp
	return arr
}

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

func (h *maxHeap) pop() int {
	// take the top of the heap
	ele := h.elements[0]
	// take the last element and make it the top of the heap
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	// rearrange the heap, given the index to start with
	h.rearrange(0)
	return ele
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

func sort(arr []int) {
	h := maxHeap{
		elements: []int{},
	}

	for i := 0; i < len(arr); i++ {
		h.push(arr[i])
	}

}

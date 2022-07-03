package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	leftCounter := calculateLeft(arr)
	rightCounter := calculateRight(arr)

	output := make([]int, len(arr))
	for i, _ := range rightCounter {
		output[i] = (leftCounter[i] + rightCounter[i]) - 1 // Minus one, to avoid counting the same subarray twice
	}

	return output
}

func calculateRight(arr []int) []int {
	// Index to the maximum element in the arr. eg: arr := []int{5,3,7} then indexMaximumElement is 2
	indexMaximumElement := len(arr) - 1

	// Sum of contiguous subarray from right side
	contiguousSubarray := make([]int, len(arr))

	// We initialize the right most index to 1. This is the maximum value that it can take looking from its right
	contiguousSubarray[indexMaximumElement] += 1

	for i := len(arr) - 2; i >= 0; i-- {

		if arr[i] >= arr[indexMaximumElement] { // if current element is bigger than the maximum element found so far
			// It adds the number of elements from where the maximum number was until the current element
			contiguousSubarray[i] += contiguousSubarray[indexMaximumElement] + (indexMaximumElement - i)
			// Update new maximum index
			indexMaximumElement = i
			continue
		}

		// Each index is a contiguous subarray then we add always 1
		contiguousSubarray[i] += 1

		if arr[i] > arr[i+1] { // If right element is bigger than the current one we add 1
			contiguousSubarray[i] += 1
		}
	}
	return contiguousSubarray
}

func calculateLeft(arr []int) []int {
	indexMaximumNumber := 0

	contiguousSubarray := make([]int, len(arr))
	contiguousSubarray[0] += 1

	for i := 1; i < len(arr); i++ {

		if arr[i] >= arr[indexMaximumNumber] {
			contiguousSubarray[i] += contiguousSubarray[indexMaximumNumber] + (i - indexMaximumNumber)
			indexMaximumNumber = i
			continue
		}

		contiguousSubarray[i] += 1
		if arr[i] > arr[i-1] {
			contiguousSubarray[i] += 1
		}
	}

	return contiguousSubarray
}

func main() {
	countSubarrays([]int{6, 4, 7})
}

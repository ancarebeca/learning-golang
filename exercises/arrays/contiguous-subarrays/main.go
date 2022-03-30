package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	leftCounter := make([]int, len(arr))

	indexMaximumNumber := 0
	contiguousSubarrays := 0
	leftCounter[0] += 1
	for i := 1; i < len(arr); i++ {
		currentElement := arr[i]
		left := i - 1
		leftCounter[i] += 1

		if currentElement >= arr[indexMaximumNumber] {
			leftCounter[i] += leftCounter[indexMaximumNumber] + (i - indexMaximumNumber - 1)
			indexMaximumNumber = i
		} else if currentElement > arr[left] {
			contiguousSubarrays += 1
			leftCounter[i] += contiguousSubarrays
		} else {
			contiguousSubarrays = 0
		}
	}

	indexMaximumNumber = len(arr) - 1
	contiguousSubarrays = 0
	rightCounter := make([]int, len(arr))

	rightCounter[indexMaximumNumber] += 1
	for i := len(arr) - 2; i >= 0; i-- {
		currentElement := arr[i]
		right := i + 1
		rightCounter[i] += 1

		if currentElement >= arr[indexMaximumNumber] {
			rightCounter[i] += rightCounter[indexMaximumNumber] + (indexMaximumNumber - i - 1)
			indexMaximumNumber = i
		} else if currentElement > arr[right] {
			contiguousSubarrays += 1
			rightCounter[i] += contiguousSubarrays
		} else {
			contiguousSubarrays = 0
		}
	}

	output := make([]int, len(arr))

	for i, _ := range rightCounter {
		output[i] = leftCounter[i] + rightCounter[i] - 1
	}
	return output
}

func main() {
	// Call countSubarrays() with test cases here
}

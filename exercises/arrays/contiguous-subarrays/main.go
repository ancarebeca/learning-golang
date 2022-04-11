package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	leftCounter := calculateLeft(arr)
	rightCounter := calculateRight(arr)

	output := make([]int, len(arr))
	for i, _ := range rightCounter {
		output[i] = leftCounter[i] + rightCounter[i] - 1
	}

	return output
}

func calculateRight(arr []int) []int {
	indexMaximumNumber := len(arr) - 1
	contiguousSubarrays := 0
	rightCounter := make([]int, len(arr))

	rightCounter[indexMaximumNumber] += 1
	for i := len(arr) - 2; i >= 0; i-- {
		currentElement := arr[i]
		rightIndex := i + 1
		rightCounter[i] += 1

		if currentElement >= arr[indexMaximumNumber] {
			rightCounter[i] += rightCounter[indexMaximumNumber] + (indexMaximumNumber - i - 1)
			indexMaximumNumber = i
			continue
		}

		if currentElement > arr[rightIndex] {
			contiguousSubarrays += 1
			rightCounter[i] += contiguousSubarrays
			continue
		}
		contiguousSubarrays = 0
	}
	return rightCounter
}

func calculateLeft(arr []int) []int {
	indexMaximumNumber := 0
	contiguousSubarrays := 0

	leftCounter := make([]int, len(arr))
	leftCounter[0] += 1

	for i := 1; i < len(arr); i++ {
		currentElement := arr[i]
		leftIndex := i - 1
		leftCounter[i] += 1

		if currentElement >= arr[indexMaximumNumber] {
			leftCounter[i] += leftCounter[indexMaximumNumber] + (i - indexMaximumNumber - 1)
			indexMaximumNumber = i
			continue
		}

		if currentElement > arr[leftIndex] {
			contiguousSubarrays += 1
			leftCounter[i] += contiguousSubarrays
			continue
		}
		contiguousSubarrays = 0
	}

	return leftCounter
}

func main() {
	countSubarrays([]int{6, 4, 7})
}

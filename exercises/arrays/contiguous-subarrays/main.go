package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	subarraysCounter := make([]int, len(arr))

	max_number_index := 0
	for index, currentElement := range arr {
		subarraysCounter[index] += 1
		leftIndex := index - 1
		if index == 0 {
			continue
		}

		if currentElement > arr[leftIndex] {
			subarraysCounter[index] += 1
		}

		if currentElement > arr[max_number_index] {
			numSubarrays := index - max_number_index - 1
			subarraysCounter[index] += numSubarrays
			max_number_index = index
		}
	}

	max_number_index = len(arr) - 1
	for i := len(arr) - 1; i >= 0; i-- {
		currentElement := arr[i]

		if i == 0 {
			continue
		}
		leftIndex := i - 1

		if currentElement > arr[leftIndex] {
		}

		if currentElement > arr[max_number_index] {
			numSubarrays := max_number_index - i - 2
			subarraysCounter[i] += numSubarrays
			max_number_index = i
		}
	}

	return subarraysCounter
}

func main() {
	// Call countSubarrays() with test cases here
}

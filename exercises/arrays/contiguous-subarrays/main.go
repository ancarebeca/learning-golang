package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	subarraysCounter := make([]int, len(arr))

	for index, currentElement := range arr {
		subarraysCounter[index] += 1

		rightIndex := index + 1
		for {
			if rightIndex >= (len(arr)) {
				break
			}

			rightElement := arr[rightIndex]
			if currentElement < rightElement {
				break
			}
			subarraysCounter[index] += 1
			rightIndex++
		}

		leftIndex := index - 1
		for {
			if leftIndex < 0 {
				break
			}

			leftElement := arr[leftIndex]
			if currentElement < leftElement {
				break
			}
			subarraysCounter[index] += 1
			leftIndex--
		}
	}
	return subarraysCounter
}

func main() {
	// Call countSubarrays() with test cases here
}

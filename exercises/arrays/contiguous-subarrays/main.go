package main

import "fmt"

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	subarraysCounter := make([]int, len(arr))

	for index, currentElement := range arr {
		subarraysCounter[index] += 1

		rightIndex := index + 1
		leftIndex := index - 1

		checkRight := rightIndex >= (len(arr)) || currentElement < arr[rightIndex]
		checkLeft := leftIndex < 0 || currentElement < arr[leftIndex]
		for {

			if checkRight && checkLeft {
				break
			}

			if !checkRight {
				subarraysCounter[index] += 1
				rightIndex++
				checkRight = rightIndex >= (len(arr)) || currentElement < arr[rightIndex]
			}

			if !checkLeft {
				subarraysCounter[index] += 1
				leftIndex--
				checkLeft = leftIndex < 0 || currentElement < arr[leftIndex]
			}

		}
	}
	return subarraysCounter
}

func main() {
	// Call countSubarrays() with test cases here
	fmt.Println(countSubarrays([]int{3, 4, 1, 6, 2}))
}

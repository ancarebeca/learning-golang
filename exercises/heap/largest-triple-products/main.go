package main

import (
	"fmt"
	"sort"
)

func findMaxProduct(arr []int) []int {
	output := []int{}
	maxNumbers := []int{}
	maxNumbers = append(maxNumbers, arr[0])
	maxNumbers = append(maxNumbers, arr[1])

	for i := 0; i < len(arr); i++ {
		if i < 2 {
			output = append(output, -1)
			continue
		}
		maxNumbers = appendMaxNumbers(maxNumbers, arr[i])
		output = append(output, maxNumbers[0]*maxNumbers[1]*maxNumbers[2])
	}
	return output
}

func appendMaxNumbers(maxNumbers []int, value int) []int {
	maxNumbers = append(maxNumbers, value)
	sort.Ints(maxNumbers)

	if len(maxNumbers) > 3 {
		return maxNumbers[1:]
	}

	return maxNumbers
}

func main() {
	fmt.Println(findMaxProduct([]int{1, 2, 3, 4, 5}))
	fmt.Println(findMaxProduct([]int{2, 1, 2, 1, 2}))
}

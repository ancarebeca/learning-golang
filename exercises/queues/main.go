package main

import (
	"fmt"
)

type item struct {
	value int
	index int
}

func findPositions(arr []int, x int) []int {
	var queue []item

	for i, v := range arr {
		item := item{
			value: v,
			index: i + 1,
		}

		queue = append(queue, item)
	}

	var output []int
	var temp []item
	elementsToPop := x

	for i := 0; i < x; i++ {

		if len(queue) < x {
			elementsToPop = len(queue)
		}

		// Pop elementsToPop elements from the queue and place them in a temporary queue.
		temp = queue[0:elementsToPop]
		queue = queue[elementsToPop:]

		max := item{}
		var indexToRemove int
		for j, item := range temp {
			// Find the largest value (if there are multiple such elements, take the one which had been popped the earliest)
			if item.value > max.value {
				indexToRemove = j
				max = item
			}

			// Decrement popped values by 1 if they're positive
			if item.value > 0 {
				temp[j].value -= 1
			}
		}

		// Store the indexToRemove of the max value
		output = append(output, temp[indexToRemove].index)

		// Remove max from temp
		temp = append(temp[:indexToRemove], temp[indexToRemove+1:]...)

		// Add the popped element back to the queue
		queue = append(queue, temp...)
	}

	return output
}

func main() {
	fmt.Println(findPositions([]int{1, 2, 2, 3, 4, 5}, 5))
}

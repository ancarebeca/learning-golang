package main

import "fmt"

func findSignatureCounts(arr []int) []int {
	 output := make([]int, len(arr))

	for i, student := range arr {
		bookOwner := student
		currentHolder := student
		for {
			output[i] +=1
			currentHolder = arr[currentHolder-1]
			if currentHolder == bookOwner{
				break
			}
		}
	}

	return output
}

func main() {
	fmt.Println(findSignatureCounts([]int{3,2,4,1}))
}
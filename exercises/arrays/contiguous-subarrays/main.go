package main

func countSubarrays(arr []int) []int {
	if len(arr) == 1 {
		return []int{1}
	}

	output := make([]int, 8)
	var subarraysCounter []int

	for index, currentElement := range arr {
		subarraysCounter[index] +=1

		nextElement := index + 1
		if currentElement > nextElement {
			subarraysCounter[index] +=1
		}


	}
	return output
}

func main() {
	// Call countSubarrays() with test cases here
}

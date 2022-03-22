package main

func findSignatureCounts(arr []int) []int {
	var signatures []int
	var next []int
	for i := 0; i < len(arr); i++ {
		for index, student := range arr {
			// Every student signs a book
			signatures[index] += 1

			// Every student pass the book
			next[student] = arr[index-student]
		}

		// check if I should stop
	}

	return nil
}

func main() {
	// Call findSignatureCounts() with test cases here
}

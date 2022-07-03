package main

func maxScoreUsingSlidingWindow(cardPoints []int, k int) int {
	total := getTotal(cardPoints)

	min := 0
	left := 0
	windowSum := 0
	sizeWindow := len(cardPoints) - k

	for right := 0; right < len(cardPoints); right++ {
		windowSum += cardPoints[right]
		if right < sizeWindow {
			min = windowSum
			continue
		}

		windowSum -= cardPoints[left]

		if windowSum < min {
			min = windowSum
		}

		left++
	}

	return total - min
}

func getTotal(cardPoints []int) int {
	total := 0
	for i := 0; i <= len(cardPoints)-1; i++ {
		total += cardPoints[i]
	}
	return total
}

/*
     k=3 total = 22
	 1 | 2 | 3 | 4 | 5 | 6 | 1
6  | _ | _ | _ | * | * | * | *      _ are our windowSum
4  | _ | _ | * | * | * | * | _      * are the number that we didn't pick. It also represents our sliding window
8  | _ | * | * | * | * | _ | _      The idea is to track of the min windowSum because total - windowSum we'll return the maximum score
12 | * | * | * | * | _ | _ | _

*/

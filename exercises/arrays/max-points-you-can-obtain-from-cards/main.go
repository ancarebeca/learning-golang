package main

func maxScoreUsingSlidingWindow(cardPoints []int, k int) int {
	total := 0

	for i := 0; i <= len(cardPoints)-1; i++ {
		total += cardPoints[i]
	}

	min := 0
	left := 0
	windowSum := 0
	sizeWindow := len(cardPoints) - k

	for right := 0; right <= len(cardPoints)-1; right++ {
		windowSum += cardPoints[right]
		if right <= sizeWindow-1 {
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

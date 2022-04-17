package main

func getMaxSum(arr []int, k int) int {
	max := 0

	windowSum := 0

	left := 0

	for right := 0; right <= len(arr)-1; right++ {

		windowSum += arr[right]

		if right < k {
			max = windowSum
			continue
		}

		windowSum -= arr[left]

		if windowSum > max {
			max = windowSum
		}

		left++
	}
	return max
}

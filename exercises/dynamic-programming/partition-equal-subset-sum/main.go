package main

import "fmt"

/** Brute Force **
The time complexity of the above algorithm is exponential O(2^n), where ‘n’ represents the total number.
The space complexity is O(n), this memory which will be used to repository the recursion stack.
**/

func canPartitionBruteForce(arr []int) bool {
	var sum int

	// find the total sum
	for i := 0; i <= len(arr)-1; i++ {
		sum += arr[i]
	}

	// if 'sum' is an odd number, we can't have two subsets with equal sum
	if (sum % 2) != 0 {
		return false
	}

	// we are trying to find a subset of given numbers that has a total sum of ‘sum/2’.
	sum /= 2

	return canPartitionRecursive(arr, sum, 0)
}

func canPartitionRecursive(arr []int, sum, currentIndex int) bool {
	fmt.Printf("t %v, i %v\n", sum, currentIndex)

	// base check
	if sum == 0 {
		return true
	}

	if len(arr) == 0 || currentIndex >= len(arr) {
		return false
	}

	// recursive call after choosing the number at the currentIndex
	// if the number at currentIndex exceeds the sum, we shouldn't process this
	if arr[currentIndex] <= sum {
		if canPartitionRecursive(arr, sum-arr[currentIndex], currentIndex+1) {
			return true
		}
	}

	// recursive call after excluding the number at the currentIndex
	return canPartitionRecursive(arr, sum, currentIndex+1)
}

//**  Dynamic Programming **

/* Top-down approach
The algorithm has time and space complexity of O(N*S), where ‘N’ represents total numbers and ‘S’ is
the total sum of all the numbers.
*/

func canPartitionTopDown(arr []int) bool {
	var sum int

	// find the total sum
	for i := 0; i <= len(arr)-1; i++ {
		sum += arr[i]
	}

	// if 'sum' is an odd number, we can't have two subsets with equal sum
	if (sum % 2) != 0 {
		return false
	}

	// we are trying to find a subset of given numbers that has a total sum of ‘sum/2’.
	sum /= 2

	state := make(map[string]bool)
	return canPartitionRecursiveTopDown(arr, sum, 0, state)
}

func canPartitionRecursiveTopDown(arr []int, sum int, currentIndex int, states map[string]bool) bool {
	// base check
	if sum == 0 {
		return true
	}

	if len(arr) == 0 || currentIndex >= len(arr) {
		return false
	}

	currentState := fmt.Sprintf("%d-%d", sum, currentIndex)
	if value, ok := states[currentState]; ok {
		fmt.Printf("do not calculate %v", currentState)
		return value
	}

	// recursive call after choosing the number at the currentIndex
	// if the number at currentIndex exceeds the sum, we shouldn't process this
	if arr[currentIndex] <= sum {
		if canPartitionRecursiveTopDown(arr, sum-arr[currentIndex], currentIndex+1, states) {
			states[currentState] = true
			return true
		}
	}

	// recursive call after excluding the number at the currentIndex
	value := canPartitionRecursiveTopDown(arr, sum, currentIndex+1, states)
	states[currentState] = value
	return value

}

/* Bottom-up approach
The above solution has time and space complexity of O(N*S), where ‘N’ represents total numbers and ‘S’ is the total
sum of all the numbers.
*/

func canPartitionBottomUp(arr []int) bool {

	// find the total sum
	sum := 0
	for i := 0; i <= len(arr)-1; i++ {
		sum += arr[i]

	}

	// if 'sum' is a an odd number, we can't have two subsets with same total
	if sum%2 != 0 {
		return false

	}

	// we are trying to find a subset of given numbers that has a total sum of ‘sum/2’.
	sum /= 2

	dp := make([][]bool, len(arr))
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	// populate the sum=0 column, as we can always have '0' sum without including any element
	for i := 0; i <= len(arr)-1; i++ {
		dp[i][0] = true
	}

	// with only one number, we can form a subset only when the required sum is equal to its value
	for s := 1; s <= sum; s++ {
		if arr[0] == s {
			dp[0][s] = true
		} else {
			dp[0][s] = false
		}
	}

	// process all subsets for all sums
	for i := 1; i <= len(arr)-1; i++ {
		for s := 1; s <= sum; s++ {
			// if we can get the sum 's' without the number at index 'i'
			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= arr[i] { // else if we can find a subset to get the remaining sum
				dp[i][s] = dp[i-1][s-arr[i]]
			}
		}
	}

	// the bottom-right corner will have our answer.
	return dp[len(arr)-1][sum]
}

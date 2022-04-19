package main

import "fmt"

// Bottom-up Dynamic Programming
func solveKnapsackBottomUp(profits, weights []int, capacity int) int {
	// basic checks
	if capacity <= 0 || len(profits) == 0 || len(weights) != len(profits) {
		return 0
	}

	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// if we have only one weight, we will take it if it is not more than the capacity
	for c := 1; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	// process all sub-arrays for all the capacities
	for i := 1; i < len(profits); i++ {
		for c := 1; c <= capacity; c++ {
			var profit1, profit2 int
			// include the item, if it is not more than the capacity
			if weights[i] <= c {
				profit1 = profits[i] + dp[i-1][c-weights[i]]
			}
			// exclude the item
			profit2 = dp[i-1][c]
			// take maximum
			dp[i][c] = max(profit1, profit2)
		}
	}
	printSelectedElements(dp, weights, profits, capacity)

	// maximum profit will be at the bottom-right corner.
	return dp[len(profits)-1][capacity]
}

// Top-down Dynamic Programming with Memoization
func solveKnapsack(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	return knapsackRecursive(dp, profits, weights, capacity, 0)
}

func knapsackRecursive(dp [][]int, profits []int, weights []int, capacity int, currentIndex int) int {
	// base checks
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	// if we have already solved a similar problem, return the result from memory
	if dp[currentIndex][capacity] != 0 {
		return dp[currentIndex][capacity]
	}

	// recursive call after choosing the element at the currentIndex
	// if the weight of the element at currentIndex exceeds the capacity, we shouldn't process this
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackRecursive(dp, profits, weights,
			capacity-weights[currentIndex], currentIndex+1)

	}
	// recursive call after excluding the element at the currentIndex
	profit2 := knapsackRecursive(dp, profits, weights, capacity, currentIndex+1)

	dp[currentIndex][capacity] = max(profit1, profit2)
	return dp[currentIndex][capacity]
}

func printSelectedElements(dp [][]int, weights []int, profits []int, capacity int) {

	fmt.Println("Selected weights:")
	totalProfit := dp[len(weights)-1][capacity]
	for i := len(weights) - 1; i > 0; i-- {
		if totalProfit != dp[i-1][capacity] {
			fmt.Println(" ", weights[i])
			capacity -= weights[i]
			totalProfit -= profits[i]
		}
	}

	if totalProfit != 0 {
		fmt.Println(" ", weights[0])
	}
	fmt.Println("")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

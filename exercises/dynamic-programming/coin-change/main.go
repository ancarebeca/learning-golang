package main

import (
	"math"
)

func coinChangeBruteForce(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	min := math.MaxInt
	for _, c := range coins {
		target := amount - c
		totalOfCoins := coinChangeBruteForce(coins, target)
		if totalOfCoins >= 0 {
			if totalOfCoins < min {
				min = totalOfCoins
			}
		}
	}

	if min < 0 || min == math.MaxInt {
		return -1
	}

	return min + 1
}

func coinChangeMemoization(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	memoized := make([]int, amount+1)

	v := countCoin(coins, amount, memoized)
	return v
}

func countCoin(coins []int, amount int, memoized []int) int {
	if memoized[amount] != 0 {
		return memoized[amount]
	}

	minCoins := math.MaxInt32

	for i := 0; i < len(coins); i++ {
		if amount < coins[i] {
			continue
		}

		if amount == coins[i] {
			if minCoins > 1 {
				minCoins = 1
			}
			continue
		}

		// amount > coins[i]
		numCoins := countCoin(coins, amount-coins[i], memoized)
		if numCoins >= 0 && numCoins+1 < minCoins {
			minCoins = numCoins + 1
		}
	}

	if minCoins == math.MaxInt32 {
		minCoins = -1
	}

	memoized[amount] = minCoins
	return minCoins
}

// Complexity: O(amount) * len(coins)
// Space Complexity: O (amount)
func coinDynamicPrograming(coins []int, amount int) int {
	// Initialize the result array to the possible max value
	totalOfCoins := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		totalOfCoins[i] = amount + 1
	}

	// Base case
	totalOfCoins[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				totalOfCoins[a] = min(totalOfCoins[a], 1+totalOfCoins[a-c])
			}
		}
	}

	if totalOfCoins[amount] == amount+1 {
		return -1
	}

	return totalOfCoins[amount]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

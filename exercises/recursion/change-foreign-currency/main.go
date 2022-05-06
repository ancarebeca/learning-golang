package main

import "fmt"

func canGetExactChangeBruteForce(targetMoney int, denominations []int) bool {
	if targetMoney == 0 {
		return true
	}

	if targetMoney < 0 {
		return false
	}

	for _, d := range denominations {
		if canGetExactChangeBruteForce(targetMoney-d, denominations) {
			return true
		}
	}

	return false
}

func canGetExactChangeVersion2(targetMoney int, denominations []int) bool {
	if targetMoney == 0 {
		return true
	}

	if targetMoney < 0 {
		return false
	}

	for i, d := range denominations {
		if canGetExactChangeBruteForce(targetMoney-d, denominations[i:]) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canGetExactChangeVersion2(7, []int{1, 3, 4, 5}))
}

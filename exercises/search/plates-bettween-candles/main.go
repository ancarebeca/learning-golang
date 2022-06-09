package main

const plate = '*'
const candle = '|'

func findLeft(candles []int, key int) int {
	for _, c := range candles {
		if c >= key {
			return c
		}
	}
	return -1
}

func findRight(candles []int, key int) int {
	n := len(candles)
	for i := n - 1; i >= 0; i-- {
		if candles[i] <= key {
			return candles[i]
		}
	}
	return -1
}

func platesBetweenCandlesUsingBinarySearch(s string, queries [][]int) []int {
	n := len(s)
	candles := make([]int, 0)
	numOfPlates := make([]int, n)
	for i, item := range s {

		if item == candle {
			candles = append(candles, i)
		}

		if i == 0 {
			if item == plate {
				numOfPlates[i] = 1
			}
			continue
		}

		if item == candle {
			numOfPlates[i] = numOfPlates[i-1]
		} else {
			numOfPlates[i] = numOfPlates[i-1] + 1
		}
	}

	result := make([]int, len(queries))

	for i, q := range queries {
		left := findLeft(candles, q[0])
		right := findRight(candles, q[1])

		if left == -1 || right == -1 || left >= right {
			result[i] = 0
		} else {
			result[i] = numOfPlates[right] - numOfPlates[left]
		}
	}
	return result
}

func platesBetweenCandles(s string, queries [][]int) []int {
	numOfPlates := make([]int, len(s))

	prev := 0
	for i, p := range s {
		if p != plate {
			numOfPlates[i] = prev
			continue
		}
		prev++
		numOfPlates[i] += prev
	}

	leftMostCandlesCounter := make([]int, len(s))
	prevL := -1
	l := 0

	rightMostCandlesCounter := make([]int, len(s))
	r := len(s) - 1
	prevR := -1

	for l < len(s) {
		if s[l] != candle {
			leftMostCandlesCounter[l] = prevL
		} else {
			leftMostCandlesCounter[l] = l
			prevL = l
		}
		l++

		if s[r] != candle {
			rightMostCandlesCounter[r] = prevR
		} else {
			rightMostCandlesCounter[r] = r
			prevR = r
		}
		r--
	}

	plates := []int{}

	for _, q := range queries {
		x := q[0]
		nearestCandleLeft := rightMostCandlesCounter[x]

		y := q[1]
		nearestCandleRight := leftMostCandlesCounter[y]

		if nearestCandleLeft > y || nearestCandleRight < x {
			plates = append(plates, 0)
			continue
		}

		plates = append(plates, numOfPlates[nearestCandleRight]-numOfPlates[nearestCandleLeft])

	}

	return plates
}

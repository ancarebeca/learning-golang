package main

const item = '*'
const container = '|'

func counter(s string, startIndices, endIndices []int) []int {
	numOfItem := make([]int, len(s))

	prev := 0
	for i, p := range s {
		if p != item {
			numOfItem[i] = prev
			continue
		}
		prev++
		numOfItem[i] += prev
	}

	leftMostContainersCounter := make([]int, len(s))
	prevL := -1
	l := 0

	rightMostContainersCounter := make([]int, len(s))
	r := len(s) - 1
	prevR := -1

	for l < len(s) {
		if s[l] != container {
			leftMostContainersCounter[l] = prevL
		} else {
			leftMostContainersCounter[l] = l
			prevL = l
		}
		l++

		if s[r] != container {
			rightMostContainersCounter[r] = prevR
		} else {
			rightMostContainersCounter[r] = r
			prevR = r
		}
		r--
	}

	items := []int{}

	for i := 0; i < len(startIndices); i++ {
		x := startIndices[i]
		nearestContainerLeft := rightMostContainersCounter[x-1]

		y := endIndices[i]
		nearestContainerRight := leftMostContainersCounter[y-1]

		if nearestContainerLeft > y || nearestContainerRight < x {
			items = append(items, 0)
			continue
		}

		items = append(items, numOfItem[nearestContainerRight]-numOfItem[nearestContainerLeft])

	}

	return items
}

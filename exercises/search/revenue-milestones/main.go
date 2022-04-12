package main

func getMilestoneDays(revenues []int, milestones []int) []int {

	accRevenues := make([]int, len(revenues))
	accRevenues[0] = revenues[0]
	for i := 1; i <= len(revenues)-1; i++ {
		accRevenues[i] = accRevenues[i-1] + revenues[i]
	}

	var output []int
	for _, milestone := range milestones {
		output = append(output, binarySearch(accRevenues, milestone))
	}

	return output
}

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	value := -1

	for l <= r {
		m := l + (r-l)/2
		if arr[m] < target {
			l = m + 1
		}

		if target <= arr[m] {
			value = arr[m]
			r = m - 1
		}
	}

	if value == -1 {
		return -1
	}

	return l + 1
}

func main() {
	getMilestoneDays([]int{700, 800, 600, 400, 600, 700}, []int{3100, 2200, 800, 2100, 1000})
}

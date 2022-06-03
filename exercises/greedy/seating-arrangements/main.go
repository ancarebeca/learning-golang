package main

import (
	"fmt"
	"sort"
)

func minOverallAwkwardness(arr []int) int {
	sort.Ints(arr)
	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		diff = max(diff, arr[i]-arr[i-2])
	}

	return max(diff, arr[len(arr)-1]-arr[len(arr)-2])

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOverallAwkwardness([]int{5, 10, 6, 8}))
}

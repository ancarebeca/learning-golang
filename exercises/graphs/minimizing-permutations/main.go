package main

import "fmt"

func minOperations(arr []int) int {
	visited := make([]bool, len(arr))
	swap := 0
	for i := 0; i < len(arr); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true

		if arr[i] == i+1 {
			continue
		}

		nextNode := arr[i] - 1
		for !visited[nextNode] {
			visited[nextNode] = true
			nextNode = arr[nextNode] - 1
			swap++
		}
	}
	return swap
}

func main() {
	fmt.Println(minOperations([]int{3, 1, 2}))
}

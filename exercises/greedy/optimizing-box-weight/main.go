package main

import "sort"

func optimize(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	totalWeight := 0
	for i := 0; i < len(arr); i++ {
		totalWeight += arr[i]
	}

	target := totalWeight / 2
	result := []int{}
	acc := 0

	for i := 0; i < len(arr); i++ {
		result = append(result, arr[i])
		if arr[i]+acc > target {
			break
		}
		acc += arr[i]
	}
	temp := []int{}

	for i := len(result) - 1; i >= 0; i-- {
		temp = append(temp, result[i])
	}

	return temp
}

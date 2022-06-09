package main

import (
	"fmt"
	"math"
)

func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums)
	o := []int{-1, -1}

	if len(nums) == 0 {
		return o
	}

	for l < r {
		m := int(math.Floor(float64(l+r) / 2))
		if nums[m] < target {
			l = m + 1
			continue
		}

		r = m
	}

	if l < len(nums) && nums[l] == target {
		o[0] = l
	}

	l = 0
	r = len(nums)
	for l < r {
		m := int(math.Floor(float64(l+r) / 2))
		if nums[m] > target {
			r = m
			continue
		}

		l = m + 1
	}

	if r > 0 && nums[r-1] == target {
		o[1] = r - 1
	}

	return o
}

// Another solution
//func searchRange(nums []int, target int) []int {
//	l := 0
//	r := len(nums)
//	o := []int{-1, -1}
//
//	if len(nums) == 0 {
//		return o
//	}
//
//	for l < r {
//		m := int(math.Floor(float64(l+r) / 2))
//		if nums[m] < target {
//			l = m + 1
//			continue
//		}
//
//		r = m
//	}
//
//	if l >= len(nums) || nums[l] != target {
//		return o
//	}
//
//	o[0] = l
//	last := l
//
//	for i := last + 1; i < len(nums); i++ {
//		if nums[i] == target {
//			last = i
//		}
//	}
//	o[1] = last
//
//	return o
//}

func main() {
	fmt.Println(searchRange([]int{3, 3, 3}, 3))
}

//https://leetcode.com/problems/plates-between-candles/

package main

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func gridChallenge(grid []string) string {
	for i, s := range grid {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		grid[i] = string(r)
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < grid[i-1][j] {
				return "NO"
			}
		}
	}
	return "YES"
}

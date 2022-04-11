package main

import "fmt"

func minLengthSubstring(s, t string) int {

	window := make(map[string]int)
	needMap := make(map[string]int)

	for i := 0; i <= len(t)-1; i++ {
		char := string(t[i])
		needMap[char] = needMap[char] + 1
		window[char] = 0
	}

	var leftPivot, have int
	var found bool

	minimumString := len(s)
	need := len(t)

	for i := 0; i <= len(s)-1; i++ {
		char := string(s[i])
		countNeed, ok := needMap[char]
		if !ok {
			continue
		}

		window[char] = window[char] + 1

		if window[char] == countNeed {
			have++
		}

		for have == need {
			charLeft := string(s[leftPivot])

			newMinimum := i - leftPivot + 1
			if newMinimum < minimumString {
				minimumString = newMinimum
				found = true
			}

			if _, ok := window[charLeft]; ok {
				window[charLeft] -= 1
				if window[charLeft] != needMap[charLeft] {
					have--
				}
			}

			leftPivot++
		}
	}

	if !found {
		return -1
	}
	return minimumString

}

func main() {
	fmt.Println(minLengthSubstring("dcffd", "fd"))
}

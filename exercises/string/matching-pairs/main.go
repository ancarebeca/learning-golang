package main

import (
	"fmt"
	"strings"
)

func matchingPairs(s string, t string) int {
	output := 0

	if s == t {
		// case 1: any repeated characters can be swapped without need of decrease output. eg: s = "aaa" & t = "aaa"
		if containsRepeatedChars(s) {
			return len(s)
		}

		// case 2: must force a swap, which reduces output. eg. pre-swap s = 'mno', t = 'mno'; after swap s = 'mon'
		if !containsRepeatedChars(s) {
			return len(s) - 2
		}
	}

	// If s != t
	pairs := make(map[string]bool)  // matched pair and possible swaps candidates
	missed := make(map[string]bool) // non-matching pairs

	for i := 0; i < len(s); i++ {
		pairs[fmt.Sprintf("%c-%c", s[i], t[i])] = true

		if s[i] == t[i] { // Found a pair
			output++
			pairs[fmt.Sprintf("%c-%s", s[i], "true")] = true // It is a candidate to swap & it made a paired
			continue
		}

		missed[fmt.Sprintf("%c-%c", s[i], t[i])] = true
		pairs[fmt.Sprintf("%c-%s", s[i], "false")] = true // It is a candidate to swap & it didn't make a paired
	}

	// case 3: Check if the inverse make creates two new matches: eg. pre-swap s='ab', t='ba'; after swap s='ba'
	for k := range missed {
		values := strings.Split(k, "-")
		inverse := fmt.Sprintf("%s-%s", values[1], values[0])
		if pairs[inverse] {
			return output + 2 // The swap creates two new matches
		}
	}

	// case 4: Check for a swap that creates one new match: eg. pre-swap s = 'at', t = 'xa'; after swap s ='ta'
	for k := range missed {
		values := strings.Split(k, "-")
		value := fmt.Sprintf("%s-%s", values[1], "false") // We try we those values that didn't make a paired
		if pairs[value] {
			return output + 1 // The swap creates a new match
		}
	}

	// case 5 Check for a swap that breaks an existing match while creating a new one: eg. pre-swap s ='ab', t = 'aa'; after swap s = 'ba'
	for k := range missed {
		values := strings.Split(k, "-")
		value := fmt.Sprintf("%s-%s", values[1], "true") // We try we those values that matched
		if pairs[value] {
			return output // The swap didn't create any new match
		}
	}

	// case 6 check for repeated characters in 's', which can be swapped without reduces output. eg. pre-swap s = "abb", t = "axb"; after swap: s = "abb"
	if containsRepeatedChars(s) {
		return output
	}

	// From now on we have to force a swap

	// case 7 swapping an index that has a miss with an index that has a match removes a match. eg. pre-swap s ="abc", t = "abd"; after swap: s ="acb"
	if len(missed) == 1 && output >= 1 {
		return output - 1
	}

	//Case 8 when there are more than one missed pairs, whatever can be swapped without affecting counter. eg. pre-swap s ="abcu", t = "arfv"; after swap: s ="acbu"
	return output
}

func containsRepeatedChars(s string) bool {
	seen := make(map[byte]bool)
	for c := range s {
		if seen[s[c]] {
			return true
		}
		seen[s[c]] = true
	}
	return false
}

func main() {
	// Call matchingPairs() with test cases here

}

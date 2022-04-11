package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchingPairs(t *testing.T) {

	// ***** If s == t

	// case 1) any repeated characters can be swapped without need of decrease output s = "aaa" & t = "aaa"
	assert.Equal(t, 3, matchingPairs("aab", "aab"))
	assert.Equal(t, 2, matchingPairs("aa", "aa"))
	assert.Equal(t, 4, matchingPairs("aaaa", "aaaa"))

	// case 2: must force a swap, which reduces output. eg. pre-swap s='mno', t='mno'; after swap s='mon'
	assert.Equal(t, 0, matchingPairs("at", "at"))
	assert.Equal(t, 1, matchingPairs("mno", "mno"))
	assert.Equal(t, 2, matchingPairs("abcd", "abcd"))

	// ***** If s != t

	// case 3: Check if the inverse make creates two new matches: eg. pre-swap s='ab', t='ba'; after swap s='ba'
	assert.Equal(t, 2, matchingPairs("at", "ta"))
	assert.Equal(t, 2, matchingPairs("xrt", "atr"))
	assert.Equal(t, 2, matchingPairs("abt", "tca"))
	assert.Equal(t, 3, matchingPairs("abt", "tba"))
	assert.Equal(t, 3, matchingPairs("art", "atr"))
	assert.Equal(t, 3, matchingPairs("abcdc", "baccd"))
	assert.Equal(t, 4, matchingPairs("abcd", "adcb"))
	assert.Equal(t, 4, matchingPairs("abcdx", "abxcc"))
	assert.Equal(t, 5, matchingPairs("abcde", "adcbe"))

	// Case 4: Check for a swap that creates one new match: eg. pre-swap s='at', t='xa'; after swap s='ta'
	assert.Equal(t, 1, matchingPairs("at", "xa"))
	assert.Equal(t, 1, matchingPairs("zrp", "pxa"))
	assert.Equal(t, 1, matchingPairs("ax", "ya"))
	assert.Equal(t, 2, matchingPairs("ABC", "ADB"))
	assert.Equal(t, 3, matchingPairs("abxce", "abcdx"))
	assert.Equal(t, 4, matchingPairs("abcde", "axcbe"))
	assert.Equal(t, 4, matchingPairs("mnode", "mnoef"))

	// Case 5 Check for a swap that breaks an existing match while creating a new one: eg. pre-swap s='ab', t='aa'; after swap s='ba'
	assert.Equal(t, 1, matchingPairs("ab", "aa"))
	assert.Equal(t, 1, matchingPairs("acb", "aaa"))
	assert.Equal(t, 2, matchingPairs("baa", "aaa"))
	assert.Equal(t, 2, matchingPairs("bca", "baa"))
	assert.Equal(t, 2, matchingPairs("axd", "axa"))
	assert.Equal(t, 2, matchingPairs("abx", "abb"))

	// Case 6 check for repeated characters in 's', which can be swapped without reduces output. eg.  pre-swap s  ="abb", t = "axb"; after swap: s  ="abb"
	assert.Equal(t, 0, matchingPairs("aa", "bb"))
	assert.Equal(t, 0, matchingPairs("aa", "bcb"))
	assert.Equal(t, 2, matchingPairs("abb", "axb"))
	assert.Equal(t, 3, matchingPairs("abczz", "abcee"))

	// Case 7 swapping an index that has a miss with an index that has a match removes a match. eg.  pre-swap s ="abc", t = "abd"; after swap: s ="acb"
	assert.Equal(t, 0, matchingPairs("ax", "ay"))
	assert.Equal(t, 0, matchingPairs("abcd", "efgh"))
	assert.Equal(t, 1, matchingPairs("axb", "ayb"))
	assert.Equal(t, 1, matchingPairs("abc", "abd"))

	//Case 8 when there are more than one missed pairs, whatever can be swapped without affecting counter. eg. pre-swap s ="abcu", t = "arfv"; after swap: s ="acbu"
	assert.Equal(t, 0, matchingPairs("ac", "rb"))
	assert.Equal(t, 0, matchingPairs("acr", "bfe"))
	assert.Equal(t, 1, matchingPairs("abcu", "arfv"))
}

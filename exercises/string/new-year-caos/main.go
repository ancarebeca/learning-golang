package main

import (
	"strconv"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func MinimumBribes(q []int32) string {
	var totalBribes int32
	var tooChaotic bool

	for i := int32(len(q)) - 1; i >= 0; i-- {

		if q[i] == i+1 {
			continue
		}

		if q[i-1] == i+1 { // The person had to bribe one people
			totalBribes++
			q[i-1] = q[i]
			q[i] = i + 1
		} else if q[i-2] == i+1 { // The person had to bribe one people
			totalBribes += 2
			q[i-2] = q[i-1]
			q[i-1] = q[i]
			q[i] = i + 1
		} else {
			tooChaotic = true
			break
		}
	}

	if tooChaotic {
		return "Too chaotic"
	}

	return strconv.Itoa(int(totalBribes))
}

package main

import "strconv"

func SuperDigit(n string, k int32) int32 {

	if len(n) == 1 {
		value, _ := strconv.Atoi(n)
		return int32(value)
	}

	var number, sum int
	for i := 0; i < len(n); i++ {
		number, _ = strconv.Atoi(string(n[i]))
		sum += number
	}

	sum = sum * int(k)
	return SuperDigit(strconv.Itoa(sum), 1)
}

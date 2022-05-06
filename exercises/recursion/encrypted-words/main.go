package main

import (
	"fmt"
	"math"
)

func encrypt(s string) string {
	r := ""
	if len(s) == 0 {
		return ""
	}

	middle := len(s) / 2
	if len(s)%2 != 0 {
		middle = int(math.Abs(float64(middle)))
	} else {
		if middle > 0 {
			middle = middle - 1
		}
	}
	r = fmt.Sprintf("%s%s", r, string(s[middle]))
	r = fmt.Sprintf("%s%s", r, encrypt(s[:middle]))
	r = fmt.Sprintf("%s%s", r, encrypt(s[middle+1:]))

	return r
}

func main() {
	fmt.Println(encrypt("facebook"))
}

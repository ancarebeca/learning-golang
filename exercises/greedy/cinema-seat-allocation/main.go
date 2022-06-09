package main

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	seats := make(map[int]int)
	for _, v := range reservedSeats {
		seats[v[0]] |= 1 << (v[1] - 1)
	}

	output := 0
	for _, v := range seats {
		fmt.Printf("v >>1 = %b\n", v>>1)
		fmt.Printf("1 << 8 = %b\n", 1<<8)
		fmt.Printf("v >> 1 &^ (1 << 8) = %b\n", v)
		fmt.Println()
		fmt.Println()

		v = v >> 1 &^ (1 << 8)
		if v == 0 {
			output += 2
		} else if (v>>4 == 0) != (v&15 == 0) {
			output += 1
		} else if (v>>2)&^(3<<4) == 0 {
			output += 1
		}
	}
	output += 2 * (n - len(seats))
	return output
}

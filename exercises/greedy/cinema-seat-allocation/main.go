package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	seats := make(map[int]int)
	for _, v := range reservedSeats {
		seats[v[0]] |= 1 << (v[1] - 1)
	}

	output := 0
	for _, v := range seats {
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

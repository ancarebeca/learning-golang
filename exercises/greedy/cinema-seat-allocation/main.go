package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	// masking each row in reservedSeats
	seats := make(map[int]int)
	for _, v := range reservedSeats {
		seats[v[0]] |= 1 << (v[1] - 1)
	}

	//we made the assumption here that each row will contain the maximum of person family possible
	output := 2 * n
	for _, binary := range seats {
		// Mask for row can fit 2 four-person families (1000000001) = 513
		if (binary | 513) == 513 {
			continue
		} else if (binary|543) == 543 || (binary|993) == 993 || (binary|903) == 903 { // Mask for row can fit q four-person families
			// 543 => 1000011111, 993 =>1111100001, 903 -> 1110000111
			output -= 1
		} else {
			output -= 2
		}

	}
	return output
}

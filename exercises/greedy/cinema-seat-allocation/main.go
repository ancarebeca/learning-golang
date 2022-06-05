package main

import "fmt"

const cols = 10

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	seats := make(map[int]int)
	for _, v := range reservedSeats {
		row := v[0] - 1
		col := v[1] - 1
		fmt.Printf("%b\n", seats[row])
		fmt.Printf("%b\n", 1<<col)
		seats[row] |= 1 << col
		fmt.Printf("result %b\n", seats[row])
		fmt.Println("--------")
	}

	r := 0

	for _, v := range seats {
		v = v >> 1 &^ (1 << 8)
		if v == 0 {
			r += 2
		} else if (v>>4 == 0) != (v&15 == 0) {
			r += 1
		} else if (v>>2)&^(3<<4) == 0 {
			r += 1
		}
	}
	r += 2 * (n - len(seats))
	return r
}

//https://github.com/letientai299/leetcode/blob/61ee5835463dea796760386a9574d33dfd1e9b8f/go/1386.cinema-seat-allocation.go
//https://github.com/atidat/leetcode-diary/blob/0869950fa841cb9ea1d7f705b203b6bbb33829fe/leetcode/1386-M-cinema-seat-allocation/main.go
//https://github.com/yuzixun/algorithm_exercise/blob/2087972d330fcff5801c7ea55260e1cf600031e3/22/2.go
//https://github.com/Crshi/LeetCode-go/blob/1cc62d2feafe2160fd518e09dc721562f52386af/LeetCode/%E5%8F%8C%E5%91%A8%E8%B5%9B/%E5%AE%89%E6%8E%92%E7%94%B5%E5%BD%B1%E9%99%A2%E5%BA%A7%E4%BD%8D/main.go
//https://github.com/Lihe97/Leetcode-Training/blob/579c3030241f9e1e3d7870ffbd4fa7799b749804/Week_Contest/5349.%20%E5%AE%89%E6%8E%92%E7%94%B5%E5%BD%B1%E9%99%A2%E5%BA%A7%E4%BD%8D.go
//https://github.com/rhzx3519/leetcode/blob/fd4b0dd9eacf042c3ff0d41810cb45f870e20eba/go/src/leetcode/1386.%20Cinema%20Seat%20Allocation/1386.%20Cinema%20Seat%20Allocation.go

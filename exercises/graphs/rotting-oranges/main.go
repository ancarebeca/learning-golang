package main

const freshOrange = 1
const rottenOrange = 2

type rottenOrangeTacker struct {
	x    int
	y    int
	time int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	time := 0
	countFresh := 0

	var q []rottenOrangeTacker

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == freshOrange {
				countFresh++
			}

			if grid[i][j] == rottenOrange {
				q = append(q, rottenOrangeTacker{x: i, y: j, time: 0})
			}
		}
	}

	for len(q) > 0 {

		// Pop from Q
		item := q[0]
		q = q[1:]

		if item.x > 0 { // Look up
			if grid[item.x-1][item.y] == freshOrange {
				grid[item.x-1][item.y] = 2
				q = append(q, rottenOrangeTacker{x: item.x - 1, y: item.y, time: item.time + 1})
				countFresh--
			}
		}

		if item.x < rows-1 { // Look down
			if grid[item.x+1][item.y] == freshOrange {
				grid[item.x+1][item.y] = 2
				q = append(q, rottenOrangeTacker{x: item.x + 1, y: item.y, time: item.time + 1})
				countFresh--
			}
		}

		if item.y < cols-1 { // Look right
			if grid[item.x][item.y+1] == freshOrange {
				grid[item.x][item.y+1] = 2
				q = append(q, rottenOrangeTacker{x: item.x, y: item.y + 1, time: item.time + 1})
				countFresh--
			}
		}

		if item.y > 0 { // Look left
			if grid[item.x][item.y-1] == freshOrange {
				grid[item.x][item.y-1] = 2
				q = append(q, rottenOrangeTacker{x: item.x, y: item.y - 1, time: item.time + 1})
				countFresh--
			}
		}

		if countFresh == 0 {
			time = item.time
		}
	}
	if countFresh > 0 {
		return -1
	}

	return time
}

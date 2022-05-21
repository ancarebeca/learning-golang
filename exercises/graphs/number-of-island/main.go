package main

// Time complexity: O(nXm)
// Worst case it will traverse a cell 5 times. Once while processing the cell and 4 times from adjacent cell calls

// Space complexity: O(nxm)
// This is are recursive algorithm with no special data structure. So the space is determined mostly by how many recursive
// stack call are made. If the entire grid is all island, After the first island it just keep calling DFS for every other island
// until itt hits O(mXm) space

func numIslands(grid [][]byte) int {
	rowSize := len(grid)

	if rowSize == 0 {
		return 0
	}

	colSize := len(grid[0])

	noOfIslands := 0
	for r, rows := range grid {
		for c := range rows {

			if grid[r][c] == 1 {
				markIslandAsVisited(grid, r, c, rowSize, colSize)
				noOfIslands++
			}
		}
	}

	return noOfIslands
}

func markIslandAsVisited(grid [][]byte, row int, col int, rowSize, colSize int) {
	if row < 0 || row >= rowSize || col < 0 || col >= colSize || grid[row][col] != 1 {
		return
	}

	grid[row][col] = 2
	markIslandAsVisited(grid, row+1, col, rowSize, colSize)
	markIslandAsVisited(grid, row, col+1, rowSize, colSize)
	markIslandAsVisited(grid, row-1, col, rowSize, colSize)
	markIslandAsVisited(grid, row, col-1, rowSize, colSize)
}

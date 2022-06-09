package main

/*
	Time complexity: O(nXm) where m X m is the size of the matrix and we loop through each of the rows and columns.
    Space complexity: O(N) for the space required to repository the visited array and O(N) for the implicit call stack on DFS.
*/

func findCircleNum(isConnected [][]int) int {
	rowSize := len(isConnected)
	if rowSize == 0 {
		return 0
	}

	numOfProvinces := 0
	visited := make([]bool, rowSize)
	for r := 0; r < rowSize; r++ {
		if !visited[r] {
			numOfProvinces++
			dfs(isConnected, visited, r)
		}
	}
	return numOfProvinces
}

func dfs(grid [][]int, visited []bool, currentRow int) {
	visited[currentRow] = true
	for c := 0; c < len(grid[currentRow]); c++ {
		if grid[currentRow][c] == 1 && !visited[c] {
			visited[c] = true
			dfs(grid, visited, c)
		}
	}
}

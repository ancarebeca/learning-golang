package main

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}

			if i == 0 && j > 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
				continue
			}

			if j == 0 && i > 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
				continue
			}

			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[rows-1][cols-1]
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

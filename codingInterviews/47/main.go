package main

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Created by k at 2023/06/07 14:04
// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxValue(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := maxValue(grid)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/06/12 10:35
// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxValue(grid [][]int) (ans int) {
	// 计算第一行礼物最大价值
	for n := 1; n < len(grid[0]); n++ {
		grid[0][n] += grid[0][n-1]
	}
	// 计算第一列礼物最大价值
	for m := 1; m < len(grid); m++ {
		grid[m][0] += grid[m-1][0]
	}
	// 计算能从上方和左方到达的格子最大礼物价值
	for m := 1; m < len(grid); m++ {
		for n := 1; n < len(grid[0]); n++ {
			grid[m][n] = max(grid[m-1][n]+grid[m][n], grid[m][n-1]+grid[m][n])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

// @lc code=end

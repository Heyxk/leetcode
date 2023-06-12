// Created by k at 2023/06/12 10:20
// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

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

func maxSubArray(nums []int) (ans int) {
	ans, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxSum = max(nums[i], maxSum+nums[i])
		if ans < maxSum {
			ans = maxSum
		}
	}

	return ans
}

// @lc code=end

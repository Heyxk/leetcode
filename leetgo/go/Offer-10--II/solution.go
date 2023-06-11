// Created by k at 2023/06/11 16:06
// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numWays(n int) (ans int) {
	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

// @lc code=end

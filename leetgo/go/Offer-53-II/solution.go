// Created by k at 2023/06/10 21:49
// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func missingNumber(nums []int) (ans int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

// @lc code=end

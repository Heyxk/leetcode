// Created by k at 2023/06/10 21:41
// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func recur(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Right) && recur(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

// @lc code=end

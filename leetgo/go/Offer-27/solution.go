// Created by k at 2023/06/09 16:13
// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mirrorTree(root *TreeNode) (ans *TreeNode) {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}

	return root
}

// @lc code=end

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	// 输入树的节点数最小为1
	ret, level := 1, 0
	sum := root.Val
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level++ // 层数增加
		var v int
		for i, l := 0, len(queue); i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			v += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if v > sum {
			sum = v
			ret = level
		}
	}
	return ret
}

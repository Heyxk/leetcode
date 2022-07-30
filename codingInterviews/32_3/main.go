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
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	queue := []*TreeNode{root}
	right := true

	for len(queue) > 0 {
		tmp := []int{}

		for i, l := 0, len(queue); i < l; i++ {
			node := queue[0]
			queue = queue[1:] // 队列出一个元素
			if right {
				tmp = append(tmp, node.Val)

			} else {
				tmp = append([]int{node.Val}, tmp...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
		}
		ret = append(ret, tmp)
		right = !right
	}
	return ret

}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	var ret []int
	var queue []*TreeNode
	if root == nil {
		return []int{}
	}

	ret = append(ret, root.Val)
	queue = append(queue, root)

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:] // 队列出一个元素

		if root.Left != nil {
			ret = append(ret, root.Left.Val)
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			ret = append(ret, root.Right.Val)
			queue = append(queue, root.Right)

		}

	}
	return ret

}

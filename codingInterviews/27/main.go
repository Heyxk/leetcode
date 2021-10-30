package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var queue []*TreeNode

	// root.Left, root.Right = root.Right, root.Left
	queue = append(queue, root)

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		tmp.Left, tmp.Right = tmp.Right, tmp.Left
	}
	return root

}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)

}

func recur(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil || A.Val != B.Val {
		return false
	}

	return recur(A.Left, B.Right) && recur(A.Right, B.Left)
}

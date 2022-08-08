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
func kthLargest(root *TreeNode, k int) int {
	b := &s{res: root.Val, k: k}
	b.dfs(root)
	return b.res

}

type s struct {
	res int
	k   int
}

func (a *s) dfs(root *TreeNode) {
	if root == nil {
		return
	}

	a.dfs(root.Right)
	if a.k == 0 {
		return
	}
	if a.k--; a.k == 0 {
		a.res = root.Val
	}
	a.dfs(root.Left)

}

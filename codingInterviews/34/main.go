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

var tmp []int
var ret [][]int

func pathSum(root *TreeNode, target int) [][]int {
	tmp = tmp[:0] // clear gobal tmp
	tmp = append(tmp, 0)
	ret = ret[:0]
	dfs(root, target)
	return ret
}

func dfs(root *TreeNode, target int) {

	if root == nil {
		return
	}
	// if tmp[0]+root.Val > target {
	// 	return
	// }

	if tmp[0]+root.Val == target && root.Left == root.Right {
		tmp = append(tmp, root.Val)
		t := make([]int, len(tmp)-1)
		copy(t, tmp[1:])
		ret = append(ret, t)
		tmp = tmp[:len(tmp)-1]
		return
	}

	tmp[0] += root.Val
	tmp = append(tmp, root.Val)

	dfs(root.Left, target)
	dfs(root.Right, target)

	tmp[0] -= root.Val
	tmp = tmp[:len(tmp)-1]
}

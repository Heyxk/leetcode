package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int

	if root == nil {
		return [][]int{}
	}

	ret = append(ret, []int{root.Val})

	next := []*TreeNode{root}

	for len(next) > 0 {
		if tmp := que(next); len(tmp) > 0 {
			ret = append(ret, tmp)
		}
		var tmpNext []*TreeNode
		for _, v := range next {
			if v.Left != nil {
				tmpNext = append(tmpNext, v.Left)
			}
			if v.Right != nil {
				tmpNext = append(tmpNext, v.Right)
			}
		}
		next = tmpNext
	}
	return ret
}

func que(queue []*TreeNode) []int {
	var ret []int
	for _, root := range queue {
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

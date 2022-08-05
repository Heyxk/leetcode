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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		root := &TreeNode{val, root, nil}
		return root
	}

	queue := []*TreeNode{root}
	for n := 1; len(queue) > 0; {
		n++
		for i, l := 0, len(queue); i < l; i++ {
			if n == depth {
				nodeLeft := &TreeNode{val, queue[0].Left, nil}
				nodeRight := &TreeNode{val, nil, queue[0].Right}
				queue[0].Left = nodeLeft
				queue[0].Right = nodeRight
				// return root
			} else {
				if queue[0].Left != nil {
					queue = append(queue, queue[0].Left)
				}
				if queue[0].Right != nil {
					queue = append(queue, queue[0].Right)
				}
			}
			queue = queue[1:] // 弹出队列头元素
		}
		if n == depth {
			queue = nil
		}
	}
	return root
}

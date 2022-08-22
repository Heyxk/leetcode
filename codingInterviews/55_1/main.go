package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, res := []*TreeNode{root}, 0
	for len(queue) > 0 {
		tmp := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
		res++
	}
	return res
}

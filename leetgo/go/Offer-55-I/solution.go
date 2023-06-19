// Created by k at 2023/06/19 15:07
// https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxDepth(root *TreeNode) (ans int) {
	// bfs
	/*
		if root == nil {
			return 0
		}
		queue := []*TreeNode{root}

		for len(queue) != 0 {
			ans++
			n := len(queue)
			for i := 0; i < n; i++ {
				node := queue[i]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			queue = queue[n:]
		}

		return
	*/

	// dfs

	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := maxDepth(root)

	fmt.Println("\noutput:", Serialize(ans))
}

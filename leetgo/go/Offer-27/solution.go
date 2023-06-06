// Created by k at 2023/06/06 12:35
// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mirrorTree(root *TreeNode) (ans *TreeNode) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		node.Left, node.Right = node.Right, node.Left
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := mirrorTree(root)
	fmt.Println("output: " + Serialize(ans))
}

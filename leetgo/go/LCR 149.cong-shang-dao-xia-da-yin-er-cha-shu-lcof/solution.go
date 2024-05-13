// Created by k at 2024/05/13 15:36
// leetgo: dev
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func decorateRecord(root *TreeNode) (ans []int) {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := decorateRecord(root)

	fmt.Println("\noutput:", Serialize(ans))
}

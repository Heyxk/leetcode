// Created by k at 2023/06/05 16:55
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func levelOrder(root *TreeNode) (ans []int) {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		ans = append(ans, node.Val)
		queue = queue[1:]
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
	ans := levelOrder(root)
	fmt.Println("output: " + Serialize(ans))
}

// Created by k at 2023/06/09 14:32
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var tmp []int
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if len(ans)%2 == 0 {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		queue = queue[l:]
		ans = append(ans, tmp)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := levelOrder(root)

	fmt.Println("\noutput:", Serialize(ans))
}

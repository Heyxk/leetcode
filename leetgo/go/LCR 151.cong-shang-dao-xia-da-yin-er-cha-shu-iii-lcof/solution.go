// Created by k at 2024/05/14 15:16
// leetgo: dev
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func decorateRecord(root *TreeNode) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ql := len(queue)
		var lans []int
		for i := 0; i < ql; i++ {
			if len(ans)%2 == 0 {
				lans = append(lans, queue[i].Val)
			} else {
				lans = append([]int{queue[i].Val}, lans...)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[ql:]
		ans = append(ans, lans)
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

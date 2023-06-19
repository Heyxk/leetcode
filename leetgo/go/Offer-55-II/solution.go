// Created by k at 2023/06/19 16:08
// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(depth(root.Left)-depth(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
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
	ans := isBalanced(root)

	fmt.Println("\noutput:", Serialize(ans))
}

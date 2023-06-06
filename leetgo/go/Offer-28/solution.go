// Created by k at 2023/06/06 12:52
// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func recur(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Right) && recur(A.Right, B.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSymmetric(root)
	fmt.Println("output: " + Serialize(ans))
}

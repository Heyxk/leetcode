// Created by k at 2024/05/20 16:29
// leetgo: dev
// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}
func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	A := Deserialize[*TreeNode](ReadLine(stdin))
	B := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSubStructure(A, B)

	fmt.Println("\noutput:", Serialize(ans))
}

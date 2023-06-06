// Created by k at 2023/06/06 11:12
// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 此函数判断AB两个节点是否相等, 节点是否一一对应
func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

// 此函数判断AB两节点开始或者A左节点开始和B, 或者A右节点开始和B是否相等
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || recur(A.Right, B)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	A := Deserialize[*TreeNode](ReadLine(stdin))
	B := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSubStructure(A, B)
	fmt.Println("output: " + Serialize(ans))
}

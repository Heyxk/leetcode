// Created by k at 2023/06/16 11:27
// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type solution struct {
	ans int
	k   int
}

func (s *solution) dfs(node *TreeNode) {
	if node == nil || s.k == 0 {
		return
	}
	s.dfs(node.Right)
	s.k--
	if s.k == 0 {
		s.ans = node.Val
	}
	s.dfs(node.Left)
}

func kthLargest(root *TreeNode, k int) (ans int) {
	// 先右子树的中序遍历
	// 访问一个节点, k--, k==0时返回节点值
	s := &solution{-1, k}
	s.dfs(root)

	return s.ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := kthLargest(root, k)

	fmt.Println("\noutput:", Serialize(ans))
}

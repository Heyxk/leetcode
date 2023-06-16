// Created by k at 2023/06/16 08:57
// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type solution struct {
	ans    [][]int
	tmp    []int
	target int
}

func (s *solution) dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	s.tmp = append(s.tmp, root.Val)
	sum += root.Val
	if sum == s.target && root.Left == nil && root.Right == nil {
		s.ans = append(s.ans, append([]int{}, s.tmp...))
	}
	s.dfs(root.Left, sum)
	s.dfs(root.Right, sum)
	s.tmp = s.tmp[:len(s.tmp)-1]
}

func pathSum(root *TreeNode, target int) (ans [][]int) {
	s := &solution{make([][]int, 0), make([]int, 0), target}
	s.dfs(root, 0)

	return s.ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := pathSum(root, target)

	fmt.Println("\noutput:", Serialize(ans))
}

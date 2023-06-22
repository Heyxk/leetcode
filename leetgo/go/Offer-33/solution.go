// Created by k at 2023/06/22 20:32
// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, p-1)
}

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	postorder := Deserialize[[]int](ReadLine(stdin))
	ans := verifyPostorder(postorder)

	fmt.Println("\noutput:", Serialize(ans))
}

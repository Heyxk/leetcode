// Created by k at 2023/06/15 09:32
// https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	// 此节点已经使用过了, 先将该节点值改为空值(二进制值为0), 检查过该节点开始的前后左右之后再将该节点值复原
	board[i][j] = 0
	res := dfs(board, i+1, j, word, k+1) || dfs(board, i, j+1, word, k+1) || dfs(board, i-1, j, word, k+1) || dfs(board, i, j-1, word, k+1)
	board[i][j] = word[k]
	return res
}

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	board := Deserialize[[][]byte](ReadLine(stdin))
	word := Deserialize[string](ReadLine(stdin))
	ans := exist(board, word)

	fmt.Println("\noutput:", Serialize(ans))
}

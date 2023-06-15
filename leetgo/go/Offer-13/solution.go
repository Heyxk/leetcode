// Created by k at 2023/06/15 09:59
// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sum(x, y int) int {
	s := 0
	for x != 0 {
		s += x % 10
		x = x / 10
	}
	for y != 0 {
		s += y % 10
		y = y / 10
	}
	return s
}
func dfs(i, j, m, n, k int, visited *[100][100]bool) (ans int) {
	if i < 0 || i > m || j < 0 || j > n || sum(i, j) > k || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited)
}
func movingCount(m int, n int, k int) (ans int) {
	visied := &[100][100]bool{}
	return dfs(0, 0, m-1, n-1, k, visied)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := movingCount(m, n, k)

	fmt.Println("\noutput:", Serialize(ans))
}

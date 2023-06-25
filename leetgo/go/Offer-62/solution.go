// Created by k at 2023/06/25 22:31
// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lastRemaining(n int, m int) (ans int) {
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	ans := lastRemaining(n, m)

	fmt.Println("\noutput:", Serialize(ans))
}

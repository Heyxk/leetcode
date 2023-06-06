// Created by k at 2023/06/06 14:35
// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numWays(n int) (ans int) {
	a, b := 1, 2
	// 起始条件和斐波拉契数列不同
	for i := 1; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := numWays(n)
	fmt.Println("output: " + Serialize(ans))
}

// Created by k at 2023/06/20 10:23
// https://leetcode.cn/problems/qiu-12n-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumNums(n int) (ans int) {
	var recur func(n int) int
	recur = func(n int) int {
		_ = n > 1 && recur(n-1) > 0
		ans += n
		return ans
	}
	recur(n)

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := sumNums(n)

	fmt.Println("\noutput:", Serialize(ans))
}

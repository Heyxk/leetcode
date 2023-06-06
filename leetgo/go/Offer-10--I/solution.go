// Created by k at 2023/06/06 14:03
// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fib(n int) (ans int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b) % 1000000007
	}
	return a
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := fib(n)
	fmt.Println("output: " + Serialize(ans))
}

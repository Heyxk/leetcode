// Created by k at 2023/06/21 16:02
// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myPow(x float64, n int) (ans float64) {
	if x == 0 {
		return 0
	}
	ans = 1
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[float64](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := myPow(x, n)

	fmt.Println("\noutput:", Serialize(ans))
}

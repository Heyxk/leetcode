// Created by k at 2023/06/14 17:15
// https://leetcode.cn/problems/number-of-times-binary-string-is-prefix-aligned/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func numTimesAllBlue(flips []int) (ans int) {
	nmax := 0
	for i, v := range flips {
		nmax = max(v, nmax)
		if nmax == i+1 {
			ans++
		}
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	flips := Deserialize[[]int](ReadLine(stdin))
	ans := numTimesAllBlue(flips)

	fmt.Println("\noutput:", Serialize(ans))
}

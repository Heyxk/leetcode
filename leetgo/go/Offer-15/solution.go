// Created by k at 2023/06/22 20:59
// https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hammingWeight(num uint32) (ans int) {
	for num != 0 {
		ans += 1
		num &= num - 1
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := hammingWeight(n)

	fmt.Println("\noutput:", Serialize(ans))
}

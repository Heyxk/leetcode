// Created by k at 2023/06/06 17:11
// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSubArray(nums []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ret, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		sum = max(v, sum+v)
		ret = max(sum, ret)
	}
	return ret
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSubArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

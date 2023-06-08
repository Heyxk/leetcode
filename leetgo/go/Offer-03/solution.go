// Created by k at 2023/06/08 14:00
// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findRepeatNumber(nums []int) (ans int) {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findRepeatNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

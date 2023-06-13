// Created by k at 2023/06/13 15:59
// https://leetcode.cn/problems/number-of-unequal-triplets-in-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func unequalTriplets(nums []int) (ans int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i] != nums[k] && nums[i] != nums[j] && nums[j] != nums[k] {
					ans++
				}
			}
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := unequalTriplets(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

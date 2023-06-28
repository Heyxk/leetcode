// Created by k at 2023/06/28 23:52
// https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSlidingWindow(nums []int, k int) (ans []int) {
	if len(nums) == 0 || k == 0 {
		return []int{0}
	}
	queue := []int{}
	ans = make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]

		}
		queue = append(queue, nums[i])
	}
	ans[0] = queue[0]
	for i := k; i < len(nums); i++ {
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]

		}
		queue = append(queue, nums[i])
		ans[i-k+1] = queue[0]
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := maxSlidingWindow(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}

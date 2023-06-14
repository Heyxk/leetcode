// Created by k at 2023/06/14 09:10
// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func MYexchange(nums []int) (ans []int) {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]%2 == 0 {
			if nums[right]%2 != 0 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			} else {
				right--
			}
		} else {
			left++
		}
	}

	return nums
}

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 != 0 {
			left++
		}
		for left < right && nums[right]%2 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := exchange(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

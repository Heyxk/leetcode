// Created by k at 2023/06/08 15:07
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchRange(nums []int, target int) (ans []int) {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l <= len(nums)-1 && nums[l] == target {
		ans = append(ans, l)
	} else {
		ans = append(ans, -1)
	}

	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 && nums[r] == target {
		ans = append(ans, r)
	} else {
		ans = append(ans, -1)
	}
	return ans

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}

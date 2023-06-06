// Created by k at 2023/05/29 14:34
// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {

	ans = helper(nums, target) - helper(nums, target-1)
	return
}
func helper(nums []int, tar int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] <= tar {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)
	fmt.Println("output: " + Serialize(ans))
}

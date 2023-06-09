// Created by k at 2023/06/08 17:45
// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMin(numbers []int) (ans int) {
	findMinNum := func(nums []int, start, end int) int {
		min := nums[start]
		for start <= end {
			if nums[start] < min {
				min = nums[start]
			}
			start++
		}
		return min
	}
	l, r := 0, len(numbers)-1
	for l <= r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			return findMinNum(numbers, l, r)
		}
	}
	return numbers[l]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numbers := Deserialize[[]int](ReadLine(stdin))
	ans := minArray(numbers)

	fmt.Println("\noutput:", Serialize(ans))
}

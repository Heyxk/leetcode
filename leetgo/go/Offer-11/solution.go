// Created by k at 2023/06/05 12:49
// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minArray(numbers []int) int {
	findmin := func(nums []int, start, end int) int {
		ans := nums[start]
		for ; start <= end; start++ {
			if nums[start] > ans {
				return ans
			}
			ans = nums[start]
		}
		return ans
	}
	l, r := 0, len(numbers)-1
	for l != r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			return findmin(numbers, l, r)
		}
	}
	return numbers[l]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numbers := Deserialize[[]int](ReadLine(stdin))
	ans := minArray(numbers)
	fmt.Println("output: " + Serialize(ans))
}

// Created by k at 2024/05/15 15:39
// leetgo: dev
// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countTarget(scores []int, target int) (ans int) {

	return helper(scores, target) - helper(scores, target-1)
}
func helper(scores []int, target int) int {
	l, r := 0, len(scores)-1
	for l <= r {
		mid := l + (r-l)/2
		if scores[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	scores := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := countTarget(scores, target)

	fmt.Println("\noutput:", Serialize(ans))
}

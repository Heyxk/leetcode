// Created by k at 2023/06/25 15:46
// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findContinuousSequence(target int) (ans [][]int) {
	i, j, res := 1, 2, 3
	tmp := []int{1, 2}
	for i < j {
		if res == target {
			ans = append(ans, append([]int{}, tmp...))
		}
		if res < target {
			j++
			res += j
			tmp = append(tmp, j)
		} else {
			res -= i
			i++
			tmp = tmp[1:]
		}
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	ans := findContinuousSequence(target)

	fmt.Println("\noutput:", Serialize(ans))
}

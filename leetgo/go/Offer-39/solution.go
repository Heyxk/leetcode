// Created by k at 2023/06/24 22:59
// https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func majorityElement(nums []int) (ans int) {
	votes := 0
	for _, v := range nums {
		if votes == 0 {
			ans = v
		}
		if v == ans {
			votes += 1
		} else {
			votes += -1
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := majorityElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

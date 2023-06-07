// Created by k at 2023/06/07 14:44
// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func translateNum(num int) (ans int) {
	// 初始, 0个数字和1个数字的结果
	a := 1 // 0个数字
	b := 1 // 1个数字
	s := fmt.Sprintf("%d", num)
	// 注意i的范围
	for i := 2; i <= len(s); i++ {
		var c int
		if s[i-2:i] >= "10" && s[i-2:i] <= "25" {
			c = a + b
		} else {
			c = b
		}
		a = b
		b = c
	}

	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[int](ReadLine(stdin))
	ans := translateNum(num)

	fmt.Println("\noutput:", Serialize(ans))
}

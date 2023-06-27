// Created by k at 2023/06/27 14:59
// https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strToInt(str string) (ans int) {
	// 空字符串返回
	if len(str) == 0 {
		return 0
	}
	i := 0
	sign := 1
	// 跳过开头的空格
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i == len(str) {
		return 0
	}
	if str[i] == '-' {
		sign = -1
	}
	if str[i] == '-' || str[i] == '+' {
		i++
	}
	for _, c := range str[i:] {
		if c < '0' || c > '9' {
			break
		}
		if ans > 214748364 || ans == 214748364 && c > '7' {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		ans = ans*10 + int(c) - 48
	}

	return ans * sign
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	str := Deserialize[string](ReadLine(stdin))
	ans := strToInt(str)

	fmt.Println("\noutput:", Serialize(ans))
}

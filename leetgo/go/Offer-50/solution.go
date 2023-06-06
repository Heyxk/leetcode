// Created by k at 2023/06/05 15:45
// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func firstUniqChar(s string) (ans byte) {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}
	for i, v := range s {
		if m[v] == 1 {
			return s[i]
		}
	}
	return ' '
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := firstUniqChar(s)
	fmt.Println("output: " + Serialize(ans))
}

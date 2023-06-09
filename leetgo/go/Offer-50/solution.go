// Created by k at 2023/06/09 11:31
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
	for _, c := range s {
		m[c] += 1
	}
	for i, c := range s {
		if m[c] == 1 {
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

	fmt.Println("\noutput:", Serialize(ans))
}

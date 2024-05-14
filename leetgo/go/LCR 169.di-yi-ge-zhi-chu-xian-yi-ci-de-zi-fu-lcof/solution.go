// Created by k at 2024/05/14 15:49
// leetgo: dev
// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func dismantlingAction(arr string) (ans byte) {
	m := map[rune]int{}
	for _, c := range arr {
		m[c] += 1
	}
	for i, c := range arr {
		if m[c] == 1 {
			return arr[i]
		}
	}
	return ' '
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[string](ReadLine(stdin))
	ans := dismantlingAction(arr)

	fmt.Println("\noutput:", Serialize(ans))
}

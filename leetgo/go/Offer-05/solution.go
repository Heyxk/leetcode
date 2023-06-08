// Created by k at 2023/06/08 12:47
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func replaceSpace(s string) string {
	r := []rune{}
	for _, c := range s {
		if c == ' ' {
			r = append(r, []rune("%20")...)
		} else {
			r = append(r, c)
		}
	}
	return string(r)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := replaceSpace(s)

	fmt.Println("\noutput:", Serialize(ans))
}

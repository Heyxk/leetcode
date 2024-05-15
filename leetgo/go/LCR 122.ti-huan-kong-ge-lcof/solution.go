// Created by k at 2024/05/15 16:36
// leetgo: dev
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func pathEncryption(path string) string {
	r := []rune{}
	for _, c := range path {
		if c == '.' {
			r = append(r, ' ')
		} else {
			r = append(r, c)
		}
	}
	return string(r)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	path := Deserialize[string](ReadLine(stdin))
	ans := pathEncryption(path)

	fmt.Println("\noutput:", Serialize(ans))
}

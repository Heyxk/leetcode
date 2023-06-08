// Created by k at 2023/06/08 13:08
// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := reverseLeftWords(s, n)

	fmt.Println("\noutput:", Serialize(ans))
}

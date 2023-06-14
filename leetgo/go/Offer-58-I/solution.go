// Created by k at 2023/06/14 10:42
// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	ret := make([]byte, 0, len(s))
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		// 跳到单词第一个字符前的空格
		for i >= 0 && s[i] != ' ' {
			i--
		}
		ret = append(ret, []byte(s[i+1:j+1])...)
		ret = append(ret, ' ')
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	if len(ret) > 0 {
		return string(ret[:len(ret)-1])
	}
	return ""
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseWords(s)

	fmt.Println("\noutput:", Serialize(ans))
}

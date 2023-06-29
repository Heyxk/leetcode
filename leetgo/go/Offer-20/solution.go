// Created by k at 2023/06/27 17:05
// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	isNum, isDot, ise_or_E := false, false, false
	s = strings.Trim(s, " ")
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			isNum = true
		} else if s[i] == '.' {
			if isDot || ise_or_E {
				return false
			}
			isDot = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if !isNum || ise_or_E {
				return false
			}
			ise_or_E = true
			isNum = false
		} else if s[i] == '-' || s[i] == '+' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	return isNum

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isNumber(s)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2024/05/16 16:36
// leetgo: dev
// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseMessage(message string) string {
	var ans string
	i := len(message) - 1
	for  i >= 0 && message[i] == ' ' {
		i--
	}
	j := i
	for i >= 0 {
		for i >= 0 && (message[i]) != ' ' {
			i--
		}
		ans += message[i+1 : j+1]
		ans += " "
		for i >= 0 && message[i] == ' ' {
			i--
		}
		j = i
	}
	if len(ans) > 0 && ans[len(ans)-1] == ' '{
		ans = ans[0:len(ans)-1]
	}
	return ans

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	message := Deserialize[string](ReadLine(stdin))
	ans := reverseMessage(message)

	fmt.Println("\noutput:", Serialize(ans))
}

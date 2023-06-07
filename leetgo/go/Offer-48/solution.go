// Created by k at 2023/06/07 15:42
// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	m := make(map[byte]int)
	max := 0
	for left, right := 0, 0; right < len(s); right++ {
		if v, ok := m[s[right]]; ok && v >= left {
			left = v + 1
		}
		m[s[right]] = right
		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}

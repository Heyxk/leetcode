// Created by k at 2023/06/22 21:01
// https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func add(a int, b int) (ans int) {
	for b != 0 { // 当进位为 0 时跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	a := Deserialize[int](ReadLine(stdin))
	b := Deserialize[int](ReadLine(stdin))
	ans := add(a, b)

	fmt.Println("\noutput:", Serialize(ans))
}

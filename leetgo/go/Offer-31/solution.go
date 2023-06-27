// Created by k at 2023/06/27 09:35
// https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func validateStackSequences(pushed []int, popped []int) bool {
	i := 0
	stack := []int{}
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	pushed := Deserialize[[]int](ReadLine(stdin))
	popped := Deserialize[[]int](ReadLine(stdin))
	ans := validateStackSequences(pushed, popped)

	fmt.Println("\noutput:", Serialize(ans))
}

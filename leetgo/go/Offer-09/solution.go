// Created by k at 2023/06/08 10:13
// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {

	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.inStack = append(c.inStack, value)
}

func (c *CQueue) DeleteHead() (ans int) {
	if len(c.outStack) == 0 {
		if len(c.inStack) == 0 {
			return -1
		}
		c.outStack = append(c.outStack, c.inStack...)
		c.inStack = c.inStack[:0]
	}
	ans = c.outStack[0]
	c.outStack = c.outStack[1:]
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "appendTail":
			methodParams := MustSplitArray(params[i])
			value := Deserialize[int](methodParams[0])
			obj.AppendTail(value)
			output = append(output, "null")
		case "deleteHead":
			ans := Serialize(obj.DeleteHead())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}

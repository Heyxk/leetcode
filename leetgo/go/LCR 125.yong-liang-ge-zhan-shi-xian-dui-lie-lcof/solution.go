// Created by k at 2024/05/13 14:40
// leetgo: dev
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
	head []int
	tail []int
}

func Constructor() CQueue {

	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.tail = append(c.tail, value)
}

func (c *CQueue) DeleteHead() (ans int) {
	if len(c.tail) == 0 && len(c.head) == 0 {
		return -1
	}
	if len(c.head) == 0 {
		// c.head = append(c.head, c.tail...)
		// c.tail = c.tail[:0]
		c.head = c.tail
		c.tail = make([]int, 0)
	}
	ans = c.head[0]
	c.head = c.head[1:]

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

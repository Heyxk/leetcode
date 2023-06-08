// Created by k at 2023/06/08 14:35
// https://leetcode.cn/problems/min-stack/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MinStack struct {
	minStack []int
	stack    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	if len(m.minStack) == 0 || m.minStack[len(m.minStack)-1] >= x {
			m.minStack = append(m.minStack, x)
	}

}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
			return
	}
	top := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if m.minStack[len(m.minStack)-1] == top {
			m.minStack = m.minStack[:len(m.minStack)-1]
	}

}

func (m *MinStack) Top() (ans int) {
	if len(m.stack) != 0 {
			ans = m.stack[len(m.stack)-1]
	}
	return
}

func (m *MinStack) GetMin() (ans int) {
	ans = m.minStack[len(m.minStack)-1]
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
		case "push":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.Push(val)
			output = append(output, "null")
		case "pop":
			obj.Pop()
			output = append(output, "null")
		case "top":
			ans := Serialize(obj.Top())
			output = append(output, ans)
		case "getMin":
			ans := Serialize(obj.GetMin())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}

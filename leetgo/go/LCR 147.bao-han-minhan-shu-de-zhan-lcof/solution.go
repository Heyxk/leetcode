// Created by k at 2024/05/13 11:14
// leetgo: dev
// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/

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
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, x)
	} else if m.minStack[len(m.minStack)-1] >= x {
		m.minStack = append(m.minStack, x)
	}

}

func (m *MinStack) Pop() {
	x := m.stack[len(m.stack)-1]
	m.stack = m.stack[0 : len(m.stack)-1]
	if x == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[0 : len(m.minStack)-1]
	}
}

func (m *MinStack) Top() (ans int) {
	ans = m.stack[len(m.stack)-1]
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
			x := Deserialize[int](methodParams[0])
			obj.Push(x)
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

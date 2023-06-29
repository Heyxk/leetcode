// Created by k at 2023/06/29 09:36
// https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MaxQueue struct {
	queue    []int
	maxStack []int
}

func Constructor() MaxQueue {

	return MaxQueue{}
}

func (m *MaxQueue) Max_value() (ans int) {
	if len(m.queue) == 0 {
		return -1
	}

	return m.maxStack[0]
}

func (m *MaxQueue) Push_back(value int) {
	// 入队列
	m.queue = append(m.queue, value)
	// 保持队首元素为最大值, 入队列时将小于value的值全部弹出再将value入队列
	for len(m.maxStack) > 0 && m.maxStack[len(m.maxStack)-1] < value {
		m.maxStack = m.maxStack[:len(m.maxStack)-1]
	}
	m.maxStack = append(m.maxStack, value)
}

func (m *MaxQueue) Pop_front() (ans int) {
	if len(m.queue) == 0 {
		return -1
	}
	ans = m.queue[0]
	m.queue = m.queue[1:]
	if ans == m.maxStack[0] {
		m.maxStack = m.maxStack[1:]
	}

	return ans
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
		case "max_value":
			ans := Serialize(obj.Max_value())
			output = append(output, ans)
		case "push_back":
			methodParams := MustSplitArray(params[i])
			value := Deserialize[int](methodParams[0])
			obj.Push_back(value)
			output = append(output, "null")
		case "pop_front":
			ans := Serialize(obj.Pop_front())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}

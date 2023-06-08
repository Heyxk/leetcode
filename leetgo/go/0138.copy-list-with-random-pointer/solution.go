// Created by k at 2023/06/08 12:28
// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func copyRandomList(head *Node) (ans *Node) {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		m[ptr] = &Node{ptr.Val, nil, nil}
	}
	for ptr := head; ptr != nil; ptr = ptr.Next {
		m[ptr].Next = m[ptr.Next]
		m[ptr].Random = m[ptr.Random]
	}
	return m[head]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[Node](ReadLine(stdin))
	ans := copyRandomList(head)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/05/12 06:53
// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) (ans *ListNode) {
	var pre *ListNode
	for head != nil{
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	ans = pre
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)
	fmt.Println("output: " + Serialize(ans))
}

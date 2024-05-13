// Created by k at 2024/05/13 14:23
// leetgo: dev
// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trainningPlan(head *ListNode) (ans *ListNode) {
	var pre *ListNode
	for head != nil {
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
	ans := trainningPlan(head)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/06/08 12:14
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
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}

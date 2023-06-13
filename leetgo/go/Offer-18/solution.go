// Created by k at 2023/06/13 09:32
// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteNode(head *ListNode, val int) (ans *ListNode) {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	prev := head
	for p := head.Next; p != nil; p = p.Next {
		if p.Val == val {
			prev.Next = p.Next
			return head
		}
		prev = p
	}

	return head
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := deleteNode(head, val)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/06/13 14:41
// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	node := &ListNode{}
	ans = node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}

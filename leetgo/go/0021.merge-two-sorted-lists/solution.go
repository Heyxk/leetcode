// Created by k at 2023/06/13 15:01
// https://leetcode.cn/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	node := &ListNode{}
	ans = node
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}
	return ans.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	list1 := Deserialize[*ListNode](ReadLine(stdin))
	list2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}

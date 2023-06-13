// Created by k at 2023/06/13 09:51
// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getKthFromEnd(head *ListNode, k int) (ans *ListNode) {
	slow, fast := head, head
	for ; fast != nil; fast = fast.Next {
		if k > 0 {
			k--
			continue
		}
		slow = slow.Next
	}
	return slow
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := getKthFromEnd(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}

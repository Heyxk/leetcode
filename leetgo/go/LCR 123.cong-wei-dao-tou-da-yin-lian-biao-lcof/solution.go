// Created by k at 2024/05/16 15:45
// leetgo: dev
// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
var ans []int

func reverseBookList(head *ListNode) []int {
	/* 1 */
	// for ; head != nil; head = head.Next {
	// 	ans = append([]int{head.Val}, ans...)
	// }

	/* 2 */
	// var pre *ListNode
	// for head != nil {
	// 	next := head.Next
	// 	head.Next = pre
	// 	pre = head
	// 	head = next
	// }
	// for ; pre != nil; pre = pre.Next {
	// 	ans = append(ans, pre.Val)
	// }
	ans = nil
	recur(head)
	return ans
}

func recur(head *ListNode) {
	if head == nil {
		return
	}
	recur(head.Next)
	ans = append(ans, head.Val)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseBookList(head)

	fmt.Println("\noutput:", Serialize(ans))
}

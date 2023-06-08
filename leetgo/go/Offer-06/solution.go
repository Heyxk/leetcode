// Created by k at 2023/06/08 11:29
// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reversePrint(head *ListNode) (ans []int) {
	for ; head != nil; head = head.Next {
		ans = append([]int{head.Val}, ans...)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reversePrint(head)

	fmt.Println("\noutput:", Serialize(ans))
}

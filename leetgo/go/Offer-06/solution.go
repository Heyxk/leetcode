// Created by k at 2023/05/09 08:47
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
	for head != nil {
		ans = append([]int{head.Val}, ans...)
		head = head.Next
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reversePrint(head)
	fmt.Println("output: " + Serialize(ans))
}

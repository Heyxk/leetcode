// Created by k at 2023/06/13 15:56
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getIntersectionNode(headA, headB *ListNode) (ans *ListNode) {
	node1, node2 := headA, headB
	for node1 != node2 {
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}

	return node1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intersectVal := Deserialize[int](ReadLine(stdin))
	listA := Deserialize[*ListNode](ReadLine(stdin))
	listB := Deserialize[*ListNode](ReadLine(stdin))
	skipA := Deserialize[int](ReadLine(stdin))
	skipB := Deserialize[int](ReadLine(stdin))
	ans := getIntersectionNode(intersectVal, listA, listB, skipA, skipB)

	fmt.Println("\noutput:", Serialize(ans))
}

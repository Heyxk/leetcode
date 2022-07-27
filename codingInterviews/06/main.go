package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ret := []int{}
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}

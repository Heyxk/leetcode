package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if nil == head {
		return []int{}
	}
	var ret []int
	// tmp := make([]int, 1)
	for ; head != nil; head = head.Next {
		tmp := []int{head.Val}
		ret = append(tmp, ret...)
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

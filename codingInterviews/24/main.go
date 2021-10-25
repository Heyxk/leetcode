package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	var ret *ListNode
	for ; head != nil; head = head.Next {
		tmp := &ListNode{head.Val, nil}
		tmp.Next= ret
		ret = tmp
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

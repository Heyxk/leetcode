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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	headptr := head
	var tmp *ListNode
	for ptr.Next != nil {
		tmp = ptr.Next
		ptr.Next = tmp.Next
		tmp.Next = headptr
		headptr = tmp
	}

	return headptr
}

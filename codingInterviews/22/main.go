package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	for i := 0; i < n-k; i++ {
		head = head.Next
	}
	return head
}

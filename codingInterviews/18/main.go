package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		head = head.Next
		return head
	}
	for ptr := head; ptr.Next != nil; ptr = ptr.Next {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
			if ptr.Next == nil {
				return head
			}
		}
	}
	return head

}

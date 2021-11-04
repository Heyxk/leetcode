package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var a []*ListNode

	for ptr := head; ptr != nil; ptr = ptr.Next {
		a = append(a, ptr)
	}

	return a[len(a)-k]

}

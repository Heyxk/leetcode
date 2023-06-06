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

func reversePrint1(head *ListNode) []int {
	ret := []int{}
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}

var i, j int
var ret []int

func help(head *ListNode) {
	if head == nil {
		ret = make([]int, i)
		return
	}
	i++
	help(head.Next)
	ret[j] = head.Val
	j++
}

func reversePrint(head *ListNode) []int {
	i, j = 0, 0
	help(head)
	return ret
}

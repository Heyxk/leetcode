package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := map[*Node]*Node{}
	ptr := head
	for ptr != nil {
		m[ptr] = &Node{ptr.Val, nil, nil}
		ptr = ptr.Next
	}

	ptr = head

	for ptr != nil {
		m[ptr].Next = m[ptr.Next]
		m[ptr].Random = m[ptr.Random]
		ptr = ptr.Next
	}
	return m[head]
}

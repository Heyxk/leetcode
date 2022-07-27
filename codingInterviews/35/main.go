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
	input := map[*Node]int{}
	// output := map[int]*Node{}
	output := []*Node{}

	retHead := &Node{head.Val, nil, nil}
	input[head] = 0
	// output[0] = retHead
	output = append(output, retHead)
	for n, inPtr, outPtr := 1, head.Next, retHead; inPtr != nil; n++ {
		input[inPtr] = n
		node := &Node{inPtr.Val, nil, nil}
		// output[n] = node
		output = append(output, node)
		outPtr.Next = node
		outPtr = outPtr.Next
		inPtr = inPtr.Next
	}
	for outPtr := retHead; outPtr != nil; outPtr = outPtr.Next {
		if head.Random == nil {
			outPtr.Random = nil

		} else {
			v := input[head.Random]
			outPtr.Random = output[v]
		}
		head = head.Next

	}
	return retHead
}

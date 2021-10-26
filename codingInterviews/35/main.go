package main

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
	var nodeList, retList []*Node
	ret := &Node{Val: head.Val, Next: nil, Random: nil}
	retHead := ret
	nodeList = append(nodeList, head)
	retList = append(retList, ret)
	for h := head.Next; h != nil; h = h.Next {
		nodeList = append(nodeList, h)
		tmp := &Node{Val: h.Val, Next: nil, Random: nil}
		retList = append(retList, tmp)
		ret.Next = tmp
		ret = ret.Next
	}
	for h, h1 := head, retHead; h != nil; h, h1 = h.Next, h1.Next {
		if h.Random == nil {
			continue
		}
		for k, v := range nodeList {
			if v == h.Random {
				h1.Random = retList[k]
			}
		}
	}
	return retHead
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

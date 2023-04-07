package main

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (cq *CQueue) AppendTail(value int) {
	cq.inStack = append(cq.inStack, value)
}

func (cq *CQueue) DeleteHead() int {
	if len(cq.outStack) == 0 {
		if len(cq.inStack) == 0 {
			return -1
		}
		// cq.outStack = cq.inStack[:] // this will make cq.outStack reference the same address with cq.inStack
		cq.outStack = append(cq.outStack, cq.inStack...)
		// copy(cq.outStack, cq.inStack) // 使用copy, 如果instack比outstack大, 那么元素不能复制完
		cq.inStack = cq.inStack[:0]
	}
	ret := cq.outStack[0]
	cq.outStack = cq.outStack[1:]

	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

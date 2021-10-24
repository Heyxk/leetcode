package main

type CQueue struct {
	stack []int
	head  int
	tail  int
}

func Constructor() CQueue {
	return CQueue{stack: make([]int, 0, 10), head: 0, tail: 0}
}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)
	this.tail++
}

func (this *CQueue) DeleteHead() int {
	if this.head == this.tail {
		return -1
	}
	ret := this.stack[this.head]
	this.head++
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {

}

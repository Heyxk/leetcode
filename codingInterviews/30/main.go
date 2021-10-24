package main

type MinStack struct {
	stack   []int
	top     int
	minList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0, 5), top: 0, minList: make([]int, 0, 5)}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) > this.top {
		this.stack[this.top] = x
	} else {
		this.stack = append(this.stack, x)
	}
	this.top++
	if len(this.minList) == 0 || x <= this.stack[this.minList[len(this.minList)-1]] {
		this.minList = append(this.minList, this.top-1)
	}
}

func (this *MinStack) Pop() {
	// 题目没有说明空栈如何处理
	if this.top == 0 {
		return
	}
	if this.minList[len(this.minList)-1] == this.top-1 {
		this.minList = this.minList[:len(this.minList)-1]
	}
	this.top--
}

func (this *MinStack) Top() int {
	return this.stack[this.top-1]

}

func (this *MinStack) Min() int {
	return this.stack[this.minList[len(this.minList)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

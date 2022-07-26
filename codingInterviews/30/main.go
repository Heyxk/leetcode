package main

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0, 10), min: make([]int, 0, 10)}
}

func (ms *MinStack) Push(x int) {
	// 维护一个最小值栈 栈顶元素最大
	if len(ms.min) == 0 || x <= ms.min[len(ms.min)-1] {
		ms.min = append(ms.min, x)
	} else {
		ms.min = append(ms.min, ms.min[len(ms.min)-1])
	}
	ms.stack = append(ms.stack, x)
}

func (ms *MinStack) Pop() {
	ms.min = ms.min[:len(ms.min)-1] // min 弹出一个最小值
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1]
	}
	panic("ms.stack is empty!")
}

func (ms *MinStack) Min() int {
	if len(ms.min) > 0 {
		return ms.min[len(ms.min)-1]
	}
	panic("ms.min is empty!")
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

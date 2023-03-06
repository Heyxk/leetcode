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
	ms.stack = append(ms.stack, x)
	if len(ms.min) == 0 || x <= ms.min[0] {
		ms.min = append([]int{x}, ms.min...)
	}
}

func (ms *MinStack) Pop() {
	x := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if x == ms.min[0] {
		ms.min = ms.min[1:]
	}

}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Min() int {
	return ms.min[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

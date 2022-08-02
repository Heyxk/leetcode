package main

type MyCircularQueue struct {
	front int
	rear  int
	nums  int
	size  int
	queue []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{0, -1, 0, k, make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.nums == this.size {
		return false
	}
	this.nums++
	this.rear = (this.rear + 1) % this.size // rear 初始化值为-1
	this.queue[this.rear] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.nums == 0 {
		return false
	}
	this.nums--
	this.front = (this.front + 1) % this.size
	return true

}

func (this *MyCircularQueue) Front() int {
	if this.nums == 0 {
		return -1
	}
	return this.queue[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.nums == 0 {
		return -1
	}
	return this.queue[this.rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.nums == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.nums == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want MyCircularQueue
	}{
		// TODO: Add test cases.
		{"case1", args{0}, MyCircularQueue{0, 0, 0, 0, []int{}}},
		{"case2", args{1}, MyCircularQueue{0, 0, 0, 1, []int{0}}},
		{"case3", args{2}, MyCircularQueue{0, 0, 0, 2, []int{0, 0}}},
		{"case4", args{3}, MyCircularQueue{0, 0, 0, 3, []int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_EnQueue(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, args{1}, false},
		{"case2", fields{0, 0, 0, 1, []int{0}}, args{1}, true},
		{"case3", fields{0, 0, 2, 2, []int{0, 1}}, args{1}, false},
		{"case3", fields{0, 0, 3, 3, []int{0, 1, 2}}, args{1}, false},
		{"case4", fields{0, 0, 1, 4, []int{0, 1, 2, 3}}, args{1}, true},
		{"case5", fields{0, 0, 4, 5, []int{0, 1, 2, 3, 4}}, args{1}, true},
		{"case6", fields{0, 0, 0, 6, []int{0, 1, 2, 3, 4, 5}}, args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.EnQueue(tt.args.value); got != tt.want {
				t.Errorf("MyCircularQueue.EnQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, false},
		{"case2", fields{0, 0, 0, 1, []int{0}}, false},
		{"case3", fields{0, 0, 2, 2, []int{0, 1}}, true},
		{"case3", fields{0, 0, 3, 3, []int{0, 1, 2}}, true},
		{"case4", fields{0, 0, 1, 4, []int{0, 1, 2, 3}}, true},
		{"case5", fields{0, 0, 0, 5, []int{0, 1, 2, 3, 4}}, false},
		{"case6", fields{0, 0, 0, 6, []int{0, 1, 2, 3, 4, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.DeQueue(); got != tt.want {
				t.Errorf("MyCircularQueue.DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Front(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, -1},
		{"case2", fields{0, 0, 0, 1, []int{0}}, -1},
		{"case3", fields{1, 1, 1, 2, []int{0, 1}}, 1},
		{"case3", fields{0, 2, 3, 3, []int{0, 1, 2}}, 0},
		{"case4", fields{2, 3, 2, 4, []int{0, 1, 2, 3}}, 2},
		{"case5", fields{3, 2, 5, 5, []int{0, 1, 2, 3, 4}}, 3},
		{"case6", fields{5, 2, 4, 6, []int{0, 1, 2, 3, 4, 5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.Front(); got != tt.want {
				t.Errorf("MyCircularQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Rear(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, -1},
		{"case2", fields{0, 0, 0, 1, []int{0}}, -1},
		{"case3", fields{1, 1, 1, 2, []int{0, 1}}, 1},
		{"case3", fields{0, 2, 3, 3, []int{0, 1, 2}}, 2},
		{"case4", fields{2, 3, 2, 4, []int{0, 1, 2, 3}}, 3},
		{"case5", fields{3, 2, 5, 5, []int{0, 1, 2, 3, 4}}, 2},
		{"case6", fields{5, 2, 4, 6, []int{0, 1, 2, 3, 4, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.Rear(); got != tt.want {
				t.Errorf("MyCircularQueue.Rear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsEmpty(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, true},
		{"case2", fields{0, 0, 0, 1, []int{0}}, true},
		{"case3", fields{1, 1, 1, 2, []int{0, 1}}, false},
		{"case3", fields{0, 2, 3, 3, []int{0, 1, 2}}, false},
		{"case4", fields{2, 3, 2, 4, []int{0, 1, 2, 3}}, false},
		{"case5", fields{3, 2, 5, 5, []int{0, 1, 2, 3, 4}}, false},
		{"case6", fields{5, 2, 4, 6, []int{0, 1, 2, 3, 4, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsFull(t *testing.T) {
	type fields struct {
		front int
		rear  int
		nums  int
		size  int
		queue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"case1", fields{0, 0, 0, 0, []int{}}, true},
		{"case2", fields{0, 0, 0, 1, []int{0}}, false},
		{"case3", fields{1, 1, 1, 2, []int{0, 1}}, false},
		{"case3", fields{0, 2, 3, 3, []int{0, 1, 2}}, true},
		{"case4", fields{2, 3, 2, 4, []int{0, 1, 2, 3}}, false},
		{"case5", fields{3, 2, 5, 5, []int{0, 1, 2, 3, 4}}, true},
		{"case6", fields{5, 2, 4, 6, []int{0, 1, 2, 3, 4, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				nums:  tt.fields.nums,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("MyCircularQueue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

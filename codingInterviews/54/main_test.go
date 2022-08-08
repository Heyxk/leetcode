package main

import "testing"

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}

	a := &TreeNode{3, nil, nil}
	b := &TreeNode{1, nil, nil}
	c := &TreeNode{4, nil, nil}
	d := &TreeNode{2, nil, nil}
	a.Left = b
	a.Right = c
	b.Right = d

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{a, 1}, 4},
		{"case2", args{a, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_kthLargest(b *testing.B) {
	type args struct {
		root *TreeNode
		k    int
	}

	a := &TreeNode{3, nil, nil}
	b1 := &TreeNode{1, nil, nil}
	c := &TreeNode{4, nil, nil}
	d := &TreeNode{2, nil, nil}
	a.Left = b1
	a.Right = c
	b1.Right = d

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{a, 1}, 4},
		{"case2", args{a, 2}, 3},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			kthLargest(tt.args.root, tt.args.k)
		}
	}
}

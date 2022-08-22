package main

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{nil}, 0},
		{"case2", args{&TreeNode{1, nil, nil}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxDepth(b *testing.B) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{nil}, 0},
		{"case2", args{&TreeNode{1, nil, nil}}, 1},
	}
	for _, tt := range tests {
		maxDepth(tt.args.root)
	}
}

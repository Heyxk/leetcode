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

func Test_s_dfs(t *testing.T) {
	type fields struct {
		res int
		k   int
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &s{
				res: tt.fields.res,
				k:   tt.fields.k,
			}
			a.dfs(tt.args.root)
		})
	}
}

package main

import "testing"

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	a := &TreeNode{3, nil, nil}
	b := &TreeNode{4, nil, nil}
	c := &TreeNode{5, nil, nil}
	d := &TreeNode{1, nil, nil}
	e := &TreeNode{2, nil, nil}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e

	aa := &TreeNode{4, nil, nil}
	bb := &TreeNode{1, nil, nil}
	aa.Left = bb

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{a, aa}, true},
		{"case2", args{nil, aa}, false},
		{"case3", args{a, nil}, false},
		{"case4", args{nil, nil}, false},
		{"case5", args{b, c}, false},
		{"case6", args{c, c}, true},
		{"case7", args{b, b}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recur(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recur(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("recur() = %v, want %v", got, tt.want)
			}
		})
	}
}

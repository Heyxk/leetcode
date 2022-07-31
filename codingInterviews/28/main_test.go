package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	a := &TreeNode{1, nil, nil}
	b := &TreeNode{2, nil, nil}
	c := &TreeNode{2, nil, nil}
	d := &TreeNode{3, nil, nil}
	e := &TreeNode{4, nil, nil}
	f := &TreeNode{4, nil, nil}
	g := &TreeNode{3, nil, nil}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g

	aa := &TreeNode{1, nil, nil}
	bb := &TreeNode{2, nil, nil}
	cc := &TreeNode{2, nil, nil}
	ff := &TreeNode{3, nil, nil}
	gg := &TreeNode{3, nil, nil}
	aa.Left = cc
	aa.Right = bb
	bb.Right = ff
	cc.Right = gg

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{a}, true},
		{"case2", args{aa}, false},
		{"case3", args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
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

func Benchmark_isSymmetric(b *testing.B) {

	type args struct {
		root *TreeNode
	}

	a := &TreeNode{1, nil, nil}
	b1 := &TreeNode{2, nil, nil}
	c := &TreeNode{2, nil, nil}
	d := &TreeNode{3, nil, nil}
	e := &TreeNode{4, nil, nil}
	f := &TreeNode{4, nil, nil}
	g := &TreeNode{3, nil, nil}
	a.Left = b1
	a.Right = c
	b1.Left = d
	b1.Right = e
	c.Left = f
	c.Right = g

	aa := &TreeNode{1, nil, nil}
	bb := &TreeNode{2, nil, nil}
	cc := &TreeNode{2, nil, nil}
	ff := &TreeNode{3, nil, nil}
	gg := &TreeNode{3, nil, nil}
	aa.Left = cc
	aa.Right = bb
	bb.Right = ff
	cc.Right = gg

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{a}, true},
		{"case2", args{aa}, false},
		{"case3", args{nil}, true},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isSymmetric(tt.args.root)
		}
	}
}

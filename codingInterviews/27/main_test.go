package main

import (
	"reflect"
	"testing"
)

func Test_mirrorTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	a := &TreeNode{4, nil, nil}
	b := &TreeNode{2, nil, nil}
	c := &TreeNode{7, nil, nil}
	d := &TreeNode{1, nil, nil}
	e := &TreeNode{3, nil, nil}
	f := &TreeNode{6, nil, nil}
	g := &TreeNode{9, nil, nil}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g

	aa := &TreeNode{4, nil, nil}
	bb := &TreeNode{2, nil, nil}
	cc := &TreeNode{7, nil, nil}
	dd := &TreeNode{1, nil, nil}
	ee := &TreeNode{3, nil, nil}
	ff := &TreeNode{6, nil, nil}
	gg := &TreeNode{9, nil, nil}
	aa.Left = cc
	aa.Right = bb
	bb.Left = ee
	bb.Right = dd
	cc.Left = gg
	cc.Right = ff

	x := &TreeNode{2, nil, nil}
	y := &TreeNode{1, nil, nil}
	z := &TreeNode{3, nil, nil}
	x.Left = y
	x.Right = z

	xx := &TreeNode{2, nil, nil}
	yy := &TreeNode{1, nil, nil}
	zz := &TreeNode{3, nil, nil}
	xx.Left = zz
	xx.Right = yy

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"case1", args{a}, aa},
		{"case2", args{nil}, nil},
		{"case3", args{z}, zz},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mirrorTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_mirrortree(b *testing.B) {
	type args struct {
		root *TreeNode
	}

	a := &TreeNode{4, nil, nil}
	b1 := &TreeNode{2, nil, nil}
	c := &TreeNode{7, nil, nil}
	d := &TreeNode{1, nil, nil}
	e := &TreeNode{3, nil, nil}
	f := &TreeNode{6, nil, nil}
	g := &TreeNode{9, nil, nil}
	a.Left = b1
	a.Right = c
	b1.Left = d
	b1.Right = e
	c.Left = f
	c.Right = g

	aa := &TreeNode{4, nil, nil}
	bb := &TreeNode{2, nil, nil}
	cc := &TreeNode{7, nil, nil}
	dd := &TreeNode{1, nil, nil}
	ee := &TreeNode{3, nil, nil}
	ff := &TreeNode{6, nil, nil}
	gg := &TreeNode{9, nil, nil}
	aa.Left = cc
	aa.Right = bb
	bb.Left = ee
	bb.Right = dd
	cc.Left = gg
	cc.Right = ff

	x := &TreeNode{2, nil, nil}
	y := &TreeNode{1, nil, nil}
	z := &TreeNode{3, nil, nil}
	x.Left = y
	x.Right = z

	xx := &TreeNode{2, nil, nil}
	yy := &TreeNode{1, nil, nil}
	zz := &TreeNode{3, nil, nil}
	xx.Left = zz
	xx.Right = yy

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"case1", args{a}, aa},
		{"case2", args{nil}, nil},
		{"case3", args{z}, zz},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mirrorTree(tt.args.root)
		}
	}
}

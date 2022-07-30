package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	a := &TreeNode{3, nil, nil}
	b := &TreeNode{9, nil, nil}
	c := &TreeNode{20, nil, nil}
	d := &TreeNode{15, nil, nil}
	e := &TreeNode{7, nil, nil}
	a.Left = b
	a.Right = c
	c.Left = d
	c.Right = e

	aa := &TreeNode{3, nil, nil}
	bb := &TreeNode{9, nil, nil}
	cc := &TreeNode{20, nil, nil}
	dd := &TreeNode{15, nil, nil}
	ee := &TreeNode{7, nil, nil}
	aa.Left = bb
	aa.Right = cc
	bb.Left = dd
	cc.Right = ee

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"case1", args{a}, [][]int{{3}, {20, 9}, {15, 7}}},
		{"case2", args{}, [][]int{}},
		{"case3", args{b}, [][]int{{9}}},
		{"case4", args{c}, [][]int{{20}, {7, 15}}},
		{"case5", args{aa}, [][]int{{3}, {20, 9}, {15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_levelOrder(b *testing.B) {

	type args struct {
		root *TreeNode
	}

	a := &TreeNode{3, nil, nil}
	b1 := &TreeNode{9, nil, nil}
	c := &TreeNode{20, nil, nil}
	d := &TreeNode{15, nil, nil}
	e := &TreeNode{7, nil, nil}
	a.Left = b1
	a.Right = c
	c.Left = d
	c.Right = e

	aa := &TreeNode{3, nil, nil}
	bb := &TreeNode{9, nil, nil}
	cc := &TreeNode{20, nil, nil}
	dd := &TreeNode{15, nil, nil}
	ee := &TreeNode{7, nil, nil}
	aa.Left = bb
	aa.Right = cc
	bb.Left = dd
	cc.Right = ee

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"case1", args{a}, [][]int{{3}, {20, 9}, {15, 7}}},
		{"case2", args{}, [][]int{}},
		{"case3", args{b1}, [][]int{{9}}},
		{"case4", args{c}, [][]int{{20}, {7, 15}}},
		{"case5", args{aa}, [][]int{{3}, {20, 9}, {15, 7}}},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			levelOrder(tt.args.root)
		}
	}
}

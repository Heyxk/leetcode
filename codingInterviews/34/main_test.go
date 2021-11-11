package main

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}

	a := &TreeNode{5, nil, nil}
	b := &TreeNode{4, nil, nil}
	c := &TreeNode{8, nil, nil}
	d := &TreeNode{11, nil, nil}
	e := &TreeNode{13, nil, nil}
	f := &TreeNode{4, nil, nil}
	g := &TreeNode{7, nil, nil}
	h := &TreeNode{2, nil, nil}
	i := &TreeNode{5, nil, nil}
	j := &TreeNode{1, nil, nil}

	a.Left = b
	a.Right = c
	b.Left = d
	c.Right = e
	c.Right = f
	d.Left = g
	d.Right = h
	f.Left = i
	f.Right = j

	aa := &TreeNode{1, nil, nil}
	bb := &TreeNode{2, nil, nil}
	cc := &TreeNode{3, nil, nil}

	aa.Left = bb
	aa.Right = cc

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"case1", args{a, 22}, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{"case2", args{aa, 5}, [][]int{}},
		{"case3", args{&TreeNode{-2, nil, &TreeNode{-3, nil,nil}}, -5}, [][]int{{-2, -3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfs(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfs(tt.args.root, tt.args.target)
		})
	}
}

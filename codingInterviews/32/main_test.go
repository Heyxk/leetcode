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

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{a}, []int{3, 9, 20, 15, 7}},
		{"case2", args{}, []int{}},
		{"case3", args{b}, []int{9}},
		{"case4", args{c}, []int{20, 15, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

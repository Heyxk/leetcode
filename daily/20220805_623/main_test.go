package main

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"case1", args{&TreeNode{4, nil, nil}, 1, 1}, &TreeNode{1, &TreeNode{4, nil, nil}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

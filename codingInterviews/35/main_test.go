package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}

	// test case
	a := &Node{1, nil, nil}
	a.Random = a
	b := &Node{2, a, a}
	c := &Node{3, b, b}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{"case1", args{b}, b},
		{"case2", args{nil}, nil},
		{"case2", args{c}, c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

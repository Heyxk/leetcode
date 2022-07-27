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
		{"case3", args{c}, c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_copyRandomList(b *testing.B) {
	type args struct {
		head *Node
	}

	// test case
	node1 := &Node{1, nil, nil}
	node1.Random = node1
	node2 := &Node{2, node1, node1}
	node3 := &Node{3, node2, node2}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{"case1", args{node2}, node2},
		{"case2", args{nil}, nil},
		{"case3", args{node3}, node3},
	}
	for _, tt := range tests {
		copyRandomList(tt.args.head)
	}
}

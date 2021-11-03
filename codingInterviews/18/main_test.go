package main

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	a := &ListNode{4, nil}
	b := &ListNode{5, nil}
	c := &ListNode{1, nil}
	d := &ListNode{9, nil}

	a.Next = b
	b.Next = c
	c.Next = d

	aa := &ListNode{4, nil}
	aa.Next = c

	bb := &ListNode{1, nil}
	cc := &ListNode{2, nil}
	dd := &ListNode{1, nil}
	bb.Next = cc

	ee := &ListNode{3, nil}
	ff := &ListNode{4, nil}
	ee.Next = ff

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{a, 5}, aa},
		{"case2", args{bb, 2}, dd},
		{"case3", args{ee, 3}, ff},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	a := &ListNode{4, nil}
	b := &ListNode{5, nil}
	c := &ListNode{1, nil}
	d := &ListNode{9, nil}

	a.Next = b
	b.Next = c
	c.Next = d

	aa := &ListNode{4, nil}
	bb := &ListNode{1, nil}
	cc := &ListNode{2, nil}
	dd := &ListNode{1, nil}
	aa.Next = bb
	bb.Next = cc
	cc.Next = dd

	ee := &ListNode{3, nil}
	ff := &ListNode{4, nil}
	ee.Next = ff

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{a, 1}, d},
		{"case2", args{aa, 2}, cc},
		{"case3", args{nil, 2}, nil},
		{"case4", args{ee, 2}, ee},
		{"case5", args{a, 3}, b},
		{"case6", args{d, 1}, d},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getKthFromEnd(b *testing.B) {
	type args struct {
		head *ListNode
		k    int
	}

	a := &ListNode{4, nil}
	b1 := &ListNode{5, nil}
	c := &ListNode{1, nil}
	d := &ListNode{9, nil}

	a.Next = b1
	b1.Next = c
	c.Next = d

	aa := &ListNode{4, nil}
	bb := &ListNode{1, nil}
	cc := &ListNode{2, nil}
	dd := &ListNode{1, nil}
	aa.Next = bb
	bb.Next = cc
	cc.Next = dd

	ee := &ListNode{3, nil}
	ff := &ListNode{4, nil}
	ee.Next = ff

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{a, 1}, d},
		{"case2", args{aa, 2}, cc},
		{"case3", args{nil, 2}, nil},
		{"case4", args{ee, 2}, ee},
		{"case5", args{a, 3}, b1},
		{"case6", args{d, 1}, d},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			getKthFromEnd(tt.args.head, tt.args.k)
		}
	}
}

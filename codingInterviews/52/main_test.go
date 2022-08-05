package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	a := &ListNode{4, nil}
	b := &ListNode{1, nil}
	c := &ListNode{8, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	aa := &ListNode{5, nil}
	bb := &ListNode{0, nil}
	cc := &ListNode{1, nil}
	aa.Next = bb
	bb.Next = cc
	cc.Next = c

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{a, aa}, c},
		{"case2", args{nil, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getIntersectionNode(b *testing.B) {

	type args struct {
		headA *ListNode
		headB *ListNode
	}

	a := &ListNode{4, nil}
	b1 := &ListNode{1, nil}
	c := &ListNode{8, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}

	a.Next = b1
	b1.Next = c
	c.Next = d
	d.Next = e

	aa := &ListNode{5, nil}
	bb := &ListNode{0, nil}
	cc := &ListNode{1, nil}
	aa.Next = bb
	bb.Next = cc
	cc.Next = c

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{a, aa}, c},
		{"case2", args{nil, nil}, nil},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			getIntersectionNode(tt.args.headA, tt.args.headB)
		}
	}
}

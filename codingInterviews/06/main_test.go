package main

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}, []int{3, 2, 1}},
		{"case2", args{nil}, []int{}},
		{"case3", args{&ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}, []int{3, 2, 2}},
		{"case4", args{&ListNode{Val: 1, Next: nil}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLeetCode(b *testing.B) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}, []int{3, 2, 1}},
		{"case2", args{nil}, []int{}},
		{"case3", args{&ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}, []int{3, 2, 2}},
		{"case4", args{&ListNode{Val: 1, Next: nil}}, []int{1}},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reversePrint(tt.args.head)
		}
	}
}

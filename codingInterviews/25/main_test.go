package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}},
		{"case2", args{nil, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
		{"case3", args{&ListNode{1, &ListNode{3, &ListNode{4, nil}}}, nil}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
		{"case4", args{nil, nil}, nil},
		{"case5", args{&ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{5, nil}}, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		{"case6", args{&ListNode{2, nil}, &ListNode{1, nil}}, &ListNode{1, &ListNode{2, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_mergeTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	}
}

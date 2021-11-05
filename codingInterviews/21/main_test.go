package main

import (
	"reflect"
	"testing"
)

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 3, 4}}, []int{3, 1, 2, 4}},
		{"case2", args{[]int{}}, []int{}},
		{"case3", args{[]int{2,16,3,5,13,1,16,1,12,18,11,8,11,11,5,1}}, []int{1, 5, 11, 11, 11, 1, 1, 13, 5, 3, 2, 16, 16, 12, 18, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

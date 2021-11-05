package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{2, 7, 11, 15}, 9}, []int{2, 7}},
		{"case2", args{[]int{10, 26, 30, 31, 47, 60}, 40}, []int{10, 30}},
		{"case3", args{[]int{10, 26, 30, 31, 47, 60}, 50}, []int{}},
		{"case4", args{[]int{10}, 10}, []int{}},
		{"case5", args{[]int{0, 10}, 10}, []int{0, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"case2", args{[]int{}}, 0},
		{"case3", args{[]int{-2, -1}}, -1},
		{"case4", args{[]int{-2, -2}}, -2},
		{"case5", args{[]int{-2}}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.prices); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

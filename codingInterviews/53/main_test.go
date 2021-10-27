package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{5, 7, 7, 8, 8, 20}, 8}, 2},
		{"case2", args{[]int{5, 7, 7, 8, 8, 20}, 6}, 0},
		{"case3", args{[]int{}, 6}, 0},
		{"case4", args{[]int{5}, 5}, 1},
		{"case5", args{[]int{5}, 4}, 0},
		{"case6", args{[]int{5, 6}, 4}, 0},
		{"case7", args{[]int{5, 6}, 6}, 1},
		{"case8", args{[]int{5, 6}, 5}, 1},
		{"case9", args{[]int{5, 6, 7, 7}, 7}, 2},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

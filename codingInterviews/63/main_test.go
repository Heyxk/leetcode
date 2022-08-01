package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"case2", args{[]int{}}, 0},
		{"case4", args{[]int{1}}, 0},
		{"case5", args{[]int{10, 11}}, 1},
		{"case6", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxProfit(b *testing.B) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"case2", args{[]int{}}, 0},
		{"case4", args{[]int{1}}, 0},
		{"case5", args{[]int{10, 11}}, 1},
		{"case6", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			maxProfit(tt.args.prices)
		}
	}
}

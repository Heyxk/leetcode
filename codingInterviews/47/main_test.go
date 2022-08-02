package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.grid); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxValue(b *testing.B) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 12},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			maxValue(tt.args.grid)
		}
	}
}

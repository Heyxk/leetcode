package main

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{2, 3, 1, 0, 2, 5, 3}}, 2},
		{"case2", args{[]int{1, 1}}, 1},
		{"case3", args{[]int{3, 4, 2, 1, 1, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findRepeatNumber(b *testing.B) {

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{2, 3, 1, 0, 2, 5, 3}}, 2},
		{"case2", args{[]int{1, 1}}, 1},
		{"case3", args{[]int{3, 4, 2, 1, 1, 0}}, 1},
	}
	for i := 0; i < b.N; i++ {

		for _, tt := range tests {
			findRepeatNumber(tt.args.nums)
		}
	}
}

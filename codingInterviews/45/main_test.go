package main

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{3, 30, 34, 5, 9}}, "3033459"},
		{"case2", args{[]int{10, 2}}, "102"},
		{"case3", args{[]int{}}, ""},
		{"case4", args{[]int{1}}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minNumber(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{3, 30, 34, 5, 9}}, "3033459"},
		{"case2", args{[]int{10, 2}}, "102"},
		{"case3", args{[]int{}}, ""},
		{"case4", args{[]int{1}}, "1"},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			minNumber(tt.args.nums)
		}
	}
}

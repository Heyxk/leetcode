package main

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 3, 4, 5}}, true},
		{"case2", args{[]int{0, 0, 1, 2, 5}}, true},
		{"case3", args{[]int{0, 0, 1, 3, 5}}, true},
		{"case4", args{[]int{0, 0, 1, 4, 5}}, true},
		{"case5", args{[]int{0, 2, 1, 4, 5}}, true},
		{"case6", args{[]int{0, 0, 1, 1, 5}}, false},
		{"case7", args{[]int{1, 2, 3, 3, 5}}, false},
		{"case8", args{[]int{1, 2, 3, 5, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isStraight(b *testing.B) {

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 3, 4, 5}}, true},
		{"case2", args{[]int{0, 0, 1, 2, 5}}, true},
		{"case3", args{[]int{0, 0, 1, 3, 5}}, true},
		{"case4", args{[]int{0, 0, 1, 4, 5}}, true},
		{"case5", args{[]int{0, 2, 1, 4, 5}}, true},
		{"case6", args{[]int{0, 0, 1, 1, 5}}, false},
		{"case7", args{[]int{1, 2, 3, 3, 5}}, false},
		{"case8", args{[]int{1, 2, 3, 5, 6}}, false},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isStraight(tt.args.nums)
		}
	}
}

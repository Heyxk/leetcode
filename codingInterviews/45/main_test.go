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

func Test_partition(t *testing.T) {
	type args struct {
		a  []int
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.a, tt.args.lo, tt.args.hi); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		a  []int
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.a, tt.args.lo, tt.args.hi)
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.a)
		})
	}
}

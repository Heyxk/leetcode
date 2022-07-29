package main

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{3, 4, 5, 1, 2}}, 1},
		{"case2", args{[]int{1, 2, 3}}, 1},
		{"case2", args{[]int{2, 3, 1}}, 1},
		{"case2", args{[]int{1}}, 1},
		{"case2", args{[]int{2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minArry(b *testing.B) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{3, 4, 5, 1, 2}}, 1},
		{"case2", args{[]int{1, 2, 3}}, 1},
		{"case2", args{[]int{2, 3, 1}}, 1},
		{"case2", args{[]int{1}}, 1},
		{"case2", args{[]int{2, 1}}, 1},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			minArray(tt.args.numbers)
		}
	}
}

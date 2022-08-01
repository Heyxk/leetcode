package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{2}, 1},
		{"case2", args{5}, 5},
		{"case3", args{0}, 0},
		{"case4", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fib(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{2}, 1},
		{"case2", args{5}, 5},
		{"case3", args{0}, 0},
		{"case4", args{1}, 1},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			fib(tt.args.n)
		}
	}
}

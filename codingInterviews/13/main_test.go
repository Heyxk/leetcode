package main

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{2, 3, 1}, 3},
		{"case2", args{3, 1, 0}, 1},
		{"case3", args{1, 1, 1}, 1},
		{"case4", args{1, 1, 0}, 1},
		{"case5", args{1, 1, 2}, 1},
		{"case6", args{1, 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

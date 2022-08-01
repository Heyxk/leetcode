package main

import "testing"

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{1}, "a"},
		{"case2", args{2}, "ab"},
		{"case3", args{3}, "abc"},
		{"case4", args{4}, "aaab"},
		{"case4", args{5}, "aaabc"},
		{"case4", args{6}, "aaaaab"},
		{"case4", args{7}, "aaaaabc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTheString(tt.args.n); got != tt.want {
				t.Errorf("generateTheString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_generateTheString(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{1}, "a"},
		{"case2", args{2}, "ab"},
		{"case3", args{3}, "abc"},
		{"case4", args{4}, "aaab"},
		{"case4", args{5}, "aaabc"},
		{"case4", args{6}, "aaaaab"},
		{"case4", args{7}, "aaaaabc"},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			generateTheString(tt.args.n)
		}
	}
}

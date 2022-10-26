package main

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"abcdefg", 2}, "cdefgab"},
		{"case2", args{"abcdefg", 0}, "abcdefg"},
		{"case3", args{"aa", 2}, "aa"},
		{"case4", args{"a", 1}, "a"},
		{"case5", args{"lrloseumgh", 6}, "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseLeftWords(b *testing.B) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"abcdefg", 2}, "cdefgab"},
		{"case2", args{"abcdefg", 0}, "abcdefg"},
		{"case3", args{"aa", 2}, "aa"},
		{"case1", args{"a", 1}, "a"},
		{"case1", args{"lrloseumgh", 6}, "umghlrlose"},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reverseLeftWords(tt.args.s, tt.args.n)
		}
	}
}

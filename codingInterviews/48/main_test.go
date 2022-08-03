package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{"abcabcdeabcdef"}, 6},
		{"case2", args{"pwwkew"}, 3},
		{"case3", args{"abcabcdeabcdef"}, 6},
		{"case4", args{"bbbbbbb"}, 1},
		{"case5", args{"abcabcbbbbbbb"}, 3},
		{"case6", args{""}, 0},
		{"case7", args{"tmmzuxt"}, 5},
		{"case8", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{"abcabcdeabcdef"}, 6},
		{"case2", args{"pwwkew"}, 3},
		{"case3", args{"abcabcdeabcdef"}, 6},
		{"case4", args{"bbbbbbb"}, 1},
		{"case5", args{"abcabcbbbbbbb"}, 3},
		{"case6", args{""}, 0},
		{"case7", args{"tmmzuxt"}, 5},
		{"case8", args{" "}, 1},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			lengthOfLongestSubstring(tt.args.s)
		}
	}
}

package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"case1", args{"abaccdeff"}, 'b'},
		{"case2", args{"aaaccddeeffb"}, 'b'},
		{"case3", args{""}, ' '},
		{"case4", args{" "}, ' '},
		{"case5", args{"abcdef"}, 'a'},
		{"case6", args{"a"}, 'a'},
		{"case7", args{"aa"}, ' '},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_firstUniqChar(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"case1", args{"abaccdeff"}, 'b'},
		{"case2", args{"aaaccddeeffb"}, 'b'},
		{"case3", args{""}, ' '},
		{"case4", args{" "}, ' '},
		{"case5", args{"abcdef"}, 'a'},
		{"case6", args{"a"}, 'a'},
		{"case7", args{"aa"}, ' '},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			firstUniqChar(tt.args.s)
		}
	}
}

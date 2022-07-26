package main

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"the sky is blue"}, "blue is sky the"},
		{"case2", args{""}, ""},
		{"case3", args{"  hello world!  "}, "world! hello"},
		{"case4", args{"s"}, "s"},
		{"case5", args{"s "}, "s"},
		{"case6", args{"s a"}, "a s"},
		{"case7", args{" "}, ""},
		{"case8", args{"  "}, ""},
		{"case9", args{"   a   b "}, "b a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = <%v>, want <%v>", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseWords(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"the sky is blue"}, "blue is sky the"},
		{"case2", args{""}, ""},
		{"case3", args{"  hello world!  "}, "world! hello"},
		{"case4", args{"s"}, "s"},
		{"case5", args{"s "}, "s"},
		{"case5", args{"s a"}, "a s"},
		{"case6", args{" "}, ""},
		{"case7", args{"  "}, ""},
		{"case9", args{"   a   b "}, "b a"},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reverseWords(tt.args.s)

		}
	}
}

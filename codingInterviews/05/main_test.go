package main

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"We are happy."}, "We%20are%20happy."},
		{"case2", args{"你 好"}, "你%20好"},
		{"case3", args{"你"}, "你"},
		{"case4", args{"a"}, "a"},
		{"case5", args{""}, ""},
		{"case6", args{" "}, "%20"},
		{"case7", args{"     "}, "%20%20%20%20%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

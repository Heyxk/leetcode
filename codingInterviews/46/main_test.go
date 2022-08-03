package main

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{12258}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_translateNum(b *testing.B) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{12258}, 5},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			translateNum(tt.args.num)
		}
	}
}

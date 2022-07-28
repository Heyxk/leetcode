package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := partition(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 3, 5, 6, 3, 5, 6}}, []int{1, 2, 3, 4, 2, 3, 4}},
		{"case2", args{[]int{40, 10, 20, 30}}, []int{4, 1, 2, 3}},
		{"case3", args{[]int{100, 100, 100, 100}}, []int{1, 1, 1, 1}},
		{"case4", args{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayRankTransform(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

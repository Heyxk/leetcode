package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"caes1", args{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5}, true},
		{"caes2", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 10}, false},
		{"caes3", args{[][]int{{10}}, 10}, true},
		{"caes4", args{[][]int{}, 10}, false},
		{"caes5", args{[][]int{{-1, 3}}, 3}, true},
		{"caes6", args{[][]int{{-1, 3}, {2, 4}, {6, 9}}, 6}, true},
		{"caes7", args{[][]int{{-1}, {2}}, 2}, true},
		{"caes8", args{[][]int{{1, 3, 5}}, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

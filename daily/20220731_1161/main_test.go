package main

import "testing"

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   7,
						Left:  &TreeNode{7, nil, nil},
						Right: &TreeNode{-8, nil, nil},
					},
					Right: &TreeNode{0, nil, nil},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxLevelSum(b *testing.B) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   7,
						Left:  &TreeNode{7, nil, nil},
						Right: &TreeNode{-8, nil, nil},
					},
					Right: &TreeNode{0, nil, nil},
				},
			},
			2,
		},
	}
	for i := 0; i > b.N; i++ {
		for _, tt := range tests {
			maxLevelSum(tt.args.root)
		}
	}
}

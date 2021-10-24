package main

import "testing"

// 单个用例
func TestShoppingOffers(t *testing.T) {
	got := shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2})
	want := 14
	if got != want {
		t.Errorf("failed, got:%d, want:%d", got, want)
	}

}

// 子测试
func TestShoppingOffers1(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := shoppingOffers([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2})
		want := 14
		if got != want {
			t.Errorf("failed, got:%d, want:%d", got, want)
		}
	})
	t.Run("case2", func(t *testing.T) {
		got := shoppingOffers([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1})
		want := 11
		if got != want {
			t.Errorf("failed, got:%d, want:%d", got, want)
		}
	})
	t.Run("case3", func(t *testing.T) {
		got := shoppingOffers([]int{1, 1, 1}, [][]int{{1, 1, 1, 0}, {2, 2, 1, 9}}, []int{1, 1, 0})
		want := 14
		if got != want {
			t.Errorf("failed, got:%d, want:%d", got, want)
		}
	})
}

// 表格驱动测试
func TestShoppingOffers2(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name    string
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{"case1", []int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}, 14},
		{"case2", []int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}, 11},
		{"case3", []int{1, 1, 1}, [][]int{{1, 1, 0, 0}, {2, 2, 1, 9}}, []int{1, 1, 0}, 0},
		{"case4", []int{1, 1, 1}, [][]int{{1, 1, 0, 0}, {2, 2, 1, 0}}, []int{1, 1, 1}, 1},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := shoppingOffers(tt.price, tt.special, tt.needs)
			if got != tt.want {
				t.Errorf("failed, got:%d, want:%d", got, tt.want)
			}
		})
	}
}

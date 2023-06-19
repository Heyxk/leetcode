// Created by k at 2023/06/19 10:32
// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j] // 控制大小根堆
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Top() int {
	return h[0]
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func getLeastNumbers(arr []int, k int) (ans []int) {
	ret := []int{}
	if k == 0 {
		return ret
	}
	h := &IntHeap{}
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if h.Top() > arr[i] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	for len(*h) > 0 {
		ret = append(ret, heap.Pop(h).(int))
	}

	return ret
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := getLeastNumbers(arr, k)

	fmt.Println("\noutput:", Serialize(ans))
}

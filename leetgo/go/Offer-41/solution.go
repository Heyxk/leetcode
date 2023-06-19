// Created by k at 2023/06/19 11:16
// https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MedianFinder struct {
	A *IntHeap
	B *IntHeap
}

type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Pop() interface{} {
	x := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

/** initialize your data structure here. */
func Constructor() MedianFinder {

	return MedianFinder{A: &IntHeap{}, B: &IntHeap{}}
}

func (m *MedianFinder) AddNum(num int) {
	// 当m==n时, 往A中插入元素
	if m.A.IntSlice.Len() == m.B.IntSlice.Len() {
		heap.Push(m.B, -num) // 大根堆采用取反保存元素
		heap.Push(m.A, -heap.Pop(m.B).(int))
	} else {
		heap.Push(m.A, num)
		heap.Push(m.B, -heap.Pop(m.A).(int)) // 注意添加到大根堆里需要取反
	}
}

func (m *MedianFinder) FindMedian() (ans float64) {
	// 使用一个小根堆A保存较大的一半数字
	// 使用一个大根堆B保存较小的一半数字
	// 假设总共N个数字, A保存m个, B保存n个, 当m==n时, 往A中插入元素(操作方式: 先将元素插入B中, 再将B中堆顶元素插入A中, 这样保证较大的元素存放于A中)
	//   当m!=n时, 往B中插入元素(操作方式: 先将元素插入A中, 再将A中堆顶元素插入B中, 这样保证较小的元素存放于B中)
	// 因为只实现了小根堆, 使用大根堆时, 将元素取反放入大根堆中, 则能实现小根堆的效果, 在弹出元素时, 再次将元素取反得到原值
	if m.A.IntSlice.Len() != m.B.IntSlice.Len() {
		return float64(m.A.IntSlice[0])
	}
	return float64(m.A.IntSlice[0]-m.B.IntSlice[0]) / 2.0 // 注意这里需要-

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "addNum":
			methodParams := MustSplitArray(params[i])
			num := Deserialize[int](methodParams[0])
			obj.AddNum(num)
			output = append(output, "null")
		case "findMedian":
			ans := Serialize(obj.FindMedian())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}

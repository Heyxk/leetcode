package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 这里决定 大小顶堆 现在是大顶堆

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
	*h = old[0 : n-1]
	return x
}

// Push 绑定push方法，插入新元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func getLeastNumbers(arr []int, k int) []int {
	ret := []int{}
	if k == 0 {
		return ret
	}
	h := &IntHeap{}
	// heap.Init(h)

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

// func main() {
// 	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
// 	heap.Init(h)                                // 将数组切片进行堆化
// 	fmt.Println(*h)                             // [0 1 3 6 2 5 7 9 8 4] 由Less方法可控制小顶堆
// 	fmt.Println(heap.Pop(h))                    // 调用pop 0 返回移除的顶部最小元素
// 	heap.Push(h, 6)                             // 调用push [1 2 3 6 4 5 7 9 8] 添加一个元素进入堆中进行堆化
// 	fmt.Println(h)                              // [1 2 3 6 4 5 7 9 8 6]
// 	for len(*h) > 0 {                           // 持续推出顶部最小元素
// 		fmt.Printf("%d \n", heap.Pop(h))
// 	}
// }

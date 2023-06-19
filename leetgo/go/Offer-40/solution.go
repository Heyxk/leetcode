// Created by k at 2023/06/19 10:32
// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j] // 控制大小根堆
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Top() int {
	return h[0]
}

func getLeastNumbers(arr []int, k int) (ans []int) {
	sort.Sort(IntHeap(arr))
	return arr[:k]

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := getLeastNumbers(arr, k)

	fmt.Println("\noutput:", Serialize(ans))
}

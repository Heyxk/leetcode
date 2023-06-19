// Created by k at 2023/06/19 09:39
// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type myintslice []int

func (m myintslice) Len() int {
	return len(m)
}
func (m myintslice) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m myintslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func isStraight(nums []int) bool {
	sort.Sort(myintslice(nums))
	joker := 0
	for i, v := range nums {
		if v == 0 {
			joker++
		} else if v == nums[(i+1)%5] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := isStraight(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

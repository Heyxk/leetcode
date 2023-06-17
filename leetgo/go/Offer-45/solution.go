// Created by k at 2023/06/17 22:47
// https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type myintslice []int

func (m myintslice) Len() int {
	return len(m)

}
func (m myintslice) Less(i, j int) bool {
	return fmt.Sprintf("%d%d", m[i], m[j]) < fmt.Sprintf("%d%d", m[j], m[i])
}
func (m myintslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func minNumber(nums []int) string {
	sort.Sort(myintslice(nums))
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(nums), " "), ""), "[]")
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}

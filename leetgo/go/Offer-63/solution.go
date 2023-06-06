// Created by k at 2023/06/06 14:48
// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) (ans int) {
	cost, profit := math.MaxInt, 0
	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}
	return profit
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)
	fmt.Println("output: " + Serialize(ans))
}

// Created by k at 2024/05/13 15:23
// leetgo: dev
// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func stockManagement(stock []int) (ans int) {
	ans = stock[0]
	for i := 0; i+1 < len(stock); i++ {
		if stock[i] > stock[i+1] {
			ans = stock[i+1]
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stock := Deserialize[[]int](ReadLine(stdin))
	ans := inventoryManagement(stock)

	fmt.Println("\noutput:", Serialize(ans))
}

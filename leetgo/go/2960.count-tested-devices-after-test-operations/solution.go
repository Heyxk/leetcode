// Created by k at 2024/05/10 17:11
// leetgo: dev
// https://leetcode.cn/problems/count-tested-devices-after-test-operations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countTestedDevices(batteryPercentages []int) (ans int) {
	for i := 0; i < len(batteryPercentages); i++ {
		// 如果不为0 则测试
		if (batteryPercentages[i] - ans) >= 1 {
			ans++
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	batteryPercentages := Deserialize[[]int](ReadLine(stdin))
	ans := countTestedDevices(batteryPercentages)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/06/25 12:42
// https://leetcode.cn/problems/jian-sheng-zi-lcof/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func cuttingRope(n int) (ans int) {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 2 {
		return int(math.Pow(3, float64(a))) * 2
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4) // 当余数为1时, 3+1 可以换成2+2, 乘积最大
	}
	return int(math.Pow(3, float64(a)))

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := cuttingRope(n)

	fmt.Println("\noutput:", Serialize(ans))
}

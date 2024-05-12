// Created by k at 2024/05/11 15:48
// leetgo: dev
// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findTargetIn2DPlants(plants [][]int, target int) bool {
	m, n := len(plants)-1, 0
	for m >= 0 && n < len(plants[0]) {
		if plants[m][n] == target {
			return true
		} else if plants[m][n] < target {
			n++
		} else {
			m--
		}
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	plants := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := findTargetIn2DPlants(plants, target)

	fmt.Println("\noutput:", Serialize(ans))
}

// Created by k at 2023/06/08 17:16
// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, 0
	for m >= 0 && n < len(matrix[0]) {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			m--
		} else {
			n++
		}
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := findNumberIn2DArray(matrix, target)

	fmt.Println("\noutput:", Serialize(ans))
}

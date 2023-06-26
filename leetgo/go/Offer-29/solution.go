// Created by k at 2023/06/26 09:40
// https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func spiralOrder(matrix [][]int) (ans []int) {
	if len(matrix) == 0 {
		return []int{}
	}
	t, b, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// 从左到右
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++ // 减掉一层
		// 到底
		if t > b {
			break
		}

		// 从上到下
		for i := t; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		// 从右到左
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[b][i])

		}
		b--
		if b < t {
			break
		}

		// 从下到上
		for i := b; i >= t; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}

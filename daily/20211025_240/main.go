package main

func searchMatrix(matrix [][]int, target int) bool {
	// 二分法
	// 二分查找x轴, 淘汰掉大于target的

	yright := len(matrix) - 1
	if yright < 0 {
		return false
	}
	yleft := 0
	for yleft < yright-1 {
		cur := (yleft + yright) / 2 // 游标
		switch {
		case matrix[cur][0] > target: // 游标值大于tatget right左移动
			yright = cur

		case matrix[cur][0] < target:
			yleft = cur
		case matrix[cur][0] == target:
			return true
		}
	}
	for i := 0; i < yright+1; i++ {
		xleft := 0
		xright := len(matrix[i]) - 1

		if xleft == xright {
			if matrix[i][0] == target {
				return true
			}
		}

		if matrix[i][xright] == target || matrix[i][xleft] == target {
			return true
		}

		for xleft < xright-1 {
			cur := (xleft + xright) / 2 // 游标
			switch {
			case matrix[i][cur] > target: // 游标值大于tatget right左移动
				xright = cur
			case matrix[i][cur] < target:
				xleft = cur
			case matrix[i][cur] == target:
				return true
			}
		}

	}
	return false

}

package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	yright := len(matrix)
	yleft := 0
LABLE0:
	for cus := (yleft + yright) / 2; yleft < yright; cus = (yleft + yright) / 2 {
		if len(matrix[cus]) == 0 {
			break LABLE0
		}
		switch {
		case matrix[cus][0] > target: // 游标值大于tatget right左移动
			if yright == cus {
				break LABLE0
			}
			yright = cus

		case matrix[cus][0] < target:
			if yleft == cus {
				break LABLE0
			}
			yleft = cus
		case matrix[cus][0] == target:
			return true
		}
	}

	for i := 0; i < yright; i++ {
		xleft := 0
		xright := len(matrix[i])
	LABLE1:
		for cus := (xleft + xright) / 2; xleft < xright; cus = (xleft + xright) / 2 {
			switch {
			case matrix[i][cus] > target: // 游标值大于tatget right左移动
				if xright == cus {
					break LABLE1
				}
				xright = cus
			case matrix[i][cus] < target:
				if xleft == cus {
					break LABLE1
				}
				xleft = cus
			case matrix[i][cus] == target:
				return true
			}
		}

	}
	return false

}

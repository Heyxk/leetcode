package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, vx := range matrix {
		for _, vy := range vx {
			if vy == target {
				return true
			} else if vy > target {
				break
			}
		}
	}
	return false

}

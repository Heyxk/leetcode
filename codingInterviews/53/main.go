package main

func search(nums []int, target int) int {
	times := 0
	for _, v := range nums {
		if v == target {
			times++
		} else if v > target {
			return times
		}
	}
	return times
}

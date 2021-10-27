package main

func missingNumber(nums []int) int {
	for k, v := range nums {
		if v != k {
			return k
		}
	}
	return len(nums)
}

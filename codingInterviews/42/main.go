package main

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < length; i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max

}

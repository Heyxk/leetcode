package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		panic("error input!")
	}
	ret, sum := nums[0], nums[0]
	for i, l := 1, len(nums); i < l; i++ {
		if sum <= 0 && nums[i] >= sum {
			sum = nums[i]
		} else {
			sum += nums[i]

		}
		if ret < sum {
			ret = sum
		}
	}
	return ret
}

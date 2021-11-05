package main

func twoSum(nums []int, target int) []int {
	j := len(nums) - 1

	for i := 0; i < j; {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}

}

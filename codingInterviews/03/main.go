package main

func findRepeatNumber(nums []int) int {
	for i, l := 0, len(nums); i < l; {
		if nums[i] == i {
			i++
			continue
		} else if nums[i] == nums[nums[i]] {
			return nums[i]
		} else {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	panic("error input")
}

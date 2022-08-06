package main

func exchange(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	for i := 1; i < length; i++ {
		if nums[i]%2 != 0 {
			nums = append([]int{nums[i]}, nums...)
			nums = append(nums[:i+1], nums[i+2:]...)
		}
	}
	return nums
}

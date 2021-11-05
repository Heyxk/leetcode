package main

func exchange(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	for i := 0; i < length; i++ {
		if nums[i]%2 != 0 {
			tmp := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append([]int{tmp}, nums...)
		}
	}
	return nums
}

package main

func search(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length
	ret := 0
LABLE:
	for cus := (left + right) / 2; left < right; cus = (left + right) / 2 {
		switch {
		case nums[cus] > target:
			if right == cus {
				break LABLE
			}
			right = cus
		case nums[cus] < target:
			if left == cus {
				break LABLE
			}
			left = cus
		case nums[cus] == target:
			for i := cus; i < length && nums[i] == target; i++ {
				ret++
			}
			for i := cus - 1; i >= 0 && nums[i] == target; i-- {
				ret++
			}
			return ret
		}
	}
	return 0
}

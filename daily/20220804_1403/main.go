package main

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func minSubsequence(nums []int) []int {
	l := len(nums)
	nums = quickSort(nums, 0, l-1)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i, tmp := 0, 0; i < l; i++ {
		tmp += nums[i]
		if sum-tmp < tmp {
			return nums[:i+1]
		}
	}
	return []int{}
}

package main

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
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

func arrayRankTransform(arr []int) []int {

	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) == 1 {
		return []int{1}
	}

	ret := []int{}
	sortArr := append([]int{}, arr...)
	sortArr = quickSort(sortArr, 0, len(arr)-1)

	index := 1
	m := map[int]int{}
	for i, v := range sortArr {
		if i == 0 {
			m[v] = index
		} else if v == sortArr[i-1] {
			continue
		} else {
			index++
			m[v] = index
		}
	}
	for _, v := range arr {
		ret = append(ret, m[v])
	}
	return ret

}

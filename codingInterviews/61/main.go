package main

func isStraight(nums []int) bool {
	var max = -1
	var min = 14
	m := make(map[int]int, 5)
	for k, v := range nums {
		if v == 0 {
			continue
		}
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = k
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
		if max-min >= 5 {
			return false
		}

	}
	return true

}

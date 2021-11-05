package main

func longestSubsequence(arr []int, difference int) int {
	m := make(map[int]int)
	var ret int
	for _, v := range arr {
		if n, ok := m[v-difference]; ok {
			m[v] = n + 1
		} else {
			m[v] = 1
		}
		if m[v] > ret {
			ret = m[v]
		}
	}
	return ret
}

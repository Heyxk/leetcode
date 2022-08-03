package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	for left, right, l := 0, 0, len(s); right < l; right++ {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}
		m[s[right]] = right
		if max < (right-left)+1 {
			max = (right - left) + 1
		}
	}
	return max
}

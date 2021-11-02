package main

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	m := make(map[byte]int)
	m[s[0]] = 0
	left, right, max := 0, 0, 0
	for right = 1; right < length; right++ {
		if v, ok := m[s[right]]; ok && v >= left {
			left = v + 1
		}
		m[s[right]] = right
		if max < (right - left + 1) {
			max = right - left + 1
		}
	}
	return max

}

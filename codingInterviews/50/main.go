package main

func firstUniqChar(s string) byte {
	slength := len(s)
	blackList := make(map[byte]bool)
LABLE:
	for i := 0; i < slength; i++ {
		if blackList[s[i]] {
			continue
		}
		for j := i + 1; j < slength; j++ {
			if s[j] == s[i] {
				blackList[s[i]] = true
				continue LABLE
			}
		}
		return s[i]
	}
	return ' '

}

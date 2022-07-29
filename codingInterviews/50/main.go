package main

func firstUniqChar(s string) byte {
	sLength := len(s)
	if sLength == 0 {
		return ' '
	}
	n := 0
	for i := 0; i < sLength; {
		if i != n && s[i] == s[n] {
			n++
			if n == sLength {
				return ' '
			}
			i = 0
		} else {
			i++
		}
	}
	return s[n]
}

package main

func reverseLeftWords(s string, n int) string {
	l := len(s)
	ret := make([]byte, 0, l)
	for i := n; i < l; i++ {
		ret = append(ret, s[i])
	}
	for i:= 0; i<n; i++ {
		ret = append(ret, s[i])
	}
	return string(ret)
}

func reverseLeftWords2(s string, n int) string {
	l := len(s)
	ret := make([]rune, 0, l)
	for i, v := range s {
		if i >= n {
			ret = append(ret, v)
		}
	}
	for i, v := range s {
		if i < n {
			ret = append(ret, v)
		}
	}
	return string(ret)
}
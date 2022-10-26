package main

func reverseLeftWords(s string, n int) string {
	l := len(s)
	ret := make([]byte, 0)
	for i := n; i < l; i++ {
		ret = append(ret, s[i])
	}
	for i:= 0; i<n; i++ {
		ret = append(ret, s[i])
	}
	return string(ret)

}

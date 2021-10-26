package main

func reverseLeftWords(s string, n int) string {
	if s == "" || n == 0 {
		return s
	}

	ret := s[n:] + s[:n]
	return ret

}

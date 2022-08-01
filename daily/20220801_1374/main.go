package main

import "strings"

func generateTheString(n int) string {
	if n == 1 {
		return "a"
	}
	s := ""

	if n%2 == 0 {
		s = strings.Join([]string{s, strings.Repeat("a", n-1)}, "")
		s = strings.Join([]string{s, "b"}, "")
	} else {
		s = strings.Join([]string{s, strings.Repeat("a", n-2)}, "")
		s = strings.Join([]string{s, "bc"}, "")
	}
	return s
}

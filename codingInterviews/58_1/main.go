package main

import "strings"

func reverseWords(s string) string {
	if len(s) == 0{
		return ""
	}
	s1 := strings.Split(s, " ")
	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	ret := ""
	for i, j := 0, len(s1); i < j; i++ {
		if s1[i] == "" {
			continue
		}
		ret += s1[i]
		ret += " "
	}
	if len(ret) > 0 {
		return ret[:len(ret)-1]
	} else {
		return ret
	}
}

package main

import (
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	tmp := strings.Split(s, " ")
	for i := 0; i < len(tmp)/2; {
		if tmp[i] == "" {
			tmp = append(tmp[:i], tmp[i+1:]...)
			continue
		}
		if tmp[len(tmp)-1-i] == "" {
			tmp = append(tmp[:len(tmp)-1-i], tmp[len(tmp)-i:]...)
			continue
		}
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
		i++
	}
	if tmp[len(tmp)/2] == "" {
		tmp = append(tmp[:len(tmp)/2], tmp[len(tmp)/2+1:]...)
	}
	return strings.Join(tmp, " ")
}

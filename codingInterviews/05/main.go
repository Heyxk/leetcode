package main

func replaceSpace(s string) string {
	ret := []rune{}
	for _, v := range s {
		if v == 32 {
			ret = append(ret, []rune("%20")...)
		} else {
			ret = append(ret, v)
		}
	}
	return string(ret)
}

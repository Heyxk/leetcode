package main

func replaceSpace(s string) string {
	sRune := make([]rune, 0, len(s))
	for _, v := range s {
		if v == 32 {
			sRune = append(sRune, []rune("%20")...)
		} else {
			sRune = append(sRune, v)
		}
	}
	return string(sRune)
}

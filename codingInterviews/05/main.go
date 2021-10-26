package main

func replaceSpace(s string) string {
	sRune := []rune(s)
	length := len(sRune)
	if length == 0 {
		return ""
	}
	r := make([]rune, 0, length*2/3)
	for i := 0; i < length; i++ {
		if sRune[i] == 32 {
			r = append(r, 37)
			r = append(r, 50)
			r = append(r, 48)
		} else {
			r = append(r, sRune[i])
		}
	}
	return string(r)
}

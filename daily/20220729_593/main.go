package main

// 判断任意三个顶点组成的三角形为等腰直角三角形, 直角边不为0, 每个三角形的直角边长度相等

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	len := -1
	valid := func(a []int, b []int, c []int) bool {
		l1 := (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
		l2 := (a[0]-c[0])*(a[0]-c[0]) + (a[1]-c[1])*(a[1]-c[1])
		l3 := (b[0]-c[0])*(b[0]-c[0]) + (b[1]-c[1])*(b[1]-c[1])
		ok := (l1 == l2 && l1+l2 == l3) || (l1 == l3 && l1+l3 == l2) || (l2 == l3 && l2+l3 == l1)
		if !ok {
			return false
		}
		if len == -1 {
			len = min(l1, l2)
		} else {
			return len != 0 && len == min(l1, l2)
		}
		return true
	}
	return valid(p1, p2, p3) && valid(p1, p2, p4) && valid(p2, p3, p4) && valid(p1, p3, p4)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

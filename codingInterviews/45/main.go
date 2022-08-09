package main

import (
	"fmt"
)

func minNumber(nums []int) string {
	quickSort(nums, 0, len(nums)-1)
	ret := ""
	for _, v := range nums {
		ret = fmt.Sprintf("%s%d", ret, v)

	}
	return ret
}

func quickSort(a []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	for i < j {

		for fmt.Sprintf("%d%d", a[j], a[l]) >= fmt.Sprintf("%d%d", a[l], a[j]) && i < j {
			j--
		}
		for fmt.Sprintf("%d%d", a[i], a[l]) <= fmt.Sprintf("%d%d", a[l], a[i]) && i < j {
			i++
		}
		a[i], a[j] = a[j], a[i]
	}
	a[i], a[l] = a[l], a[i]
	quickSort(a, l, i-1)
	quickSort(a, i+1, r)
}

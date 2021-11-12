package main

import (
	"fmt"
	"strconv"
)

func minNumber(nums []int) string {
	bubbleSort(nums)
	ret := ""
	for _, v := range nums {
		ret = fmt.Sprintf("%s%d", ret, v)

	}
	return ret
}

func bubbleSort(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			b := fmt.Sprintf("%d%d", a[i], a[j])
			c := fmt.Sprintf("%d%d", a[j], a[i])
			d, _ := strconv.Atoi(b)
			e, _ := strconv.Atoi(c)
			if d > e {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

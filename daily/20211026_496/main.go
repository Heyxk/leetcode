package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	m := make(map[int]int, 10)
	for k, v := range nums2 {
		m[v] = k
	}

	var ret []int
LABLE1:
	for _, v := range nums1 {
		for i, j := m[v], len(nums2); i < j; i++ {
			if nums2[i] > v {
				ret = append(ret, nums2[i])
				continue LABLE1
			}

		}
		ret = append(ret, -1)

	}
	return ret
}

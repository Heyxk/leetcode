package main

func minArray(numbers []int) int {
	for k, v := range numbers {
		if k-1 >= 0 && v < numbers[k-1] {
			return v
		}

	}
	return numbers[0]

}

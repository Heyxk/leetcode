package main

func minArray(numbers []int) int {
	for i, v := range numbers {
		if i == 0 {
			continue
		}
		if v < numbers[i-1] {
			return v
		}
	}
	// len(numbers) >= 1
	return numbers[0]
}

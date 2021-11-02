package main

func maxProfit(prices []int) int {

	var max int
	for i, length := 0, len(prices); i < length; i++ {
		for j := i + 1; j < length; j++ {
			if prices[i] > prices[j] {
				continue
			}
			if tmp := prices[j] - prices[i]; tmp > max {
				max = tmp
			}
		}
	}
	return max

}

package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	ret := 0
	for i, j := 0, 1; i < len(prices) && j < len(prices); {
		if prices[j] <= prices[i] {
			i = j
			j++
			continue
		}
		if (prices[j] - prices[i]) > ret {
			ret = prices[j] - prices[i]
		}
		j++
	}
	return ret
}

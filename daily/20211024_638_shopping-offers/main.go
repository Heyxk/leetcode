package main

import "fmt"

func mian() {

}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price) // 商品数量和needs长度
	specialLength := len(special)

	fmt.Println(n, specialLength)

	var specialAviable [][]int
	for i := 0; i < specialLength; i++ {
		for j := 0; j < n; j++ {
			if special[i][j] > needs[j] {
				break // 去掉大礼包中有比needs中对应货物多的
			}
		}
		specialAviable = append(specialAviable, special[i]) // 可用大礼包
	}

	// 计算每个货物单独购买的总价格 作为初始最低价格
	minPrice := 0
	for i := 0; i < n; i++ {
		minPrice = minPrice + price[i]*needs[i]
	}

	// 遍历specialAviable中的大礼包 每个大礼包分别和剩余的方式加起来计算价格
	for i, j := 0, len(specialAviable); i < j; i++ {
		//quotient remainder
		quo := needs[0] / specialAviable[i][0] // 第一个货物计算初始大礼包个数
		// rem := needs[0] % specialAviable[i][0]

		for o := 1; o < n; o++ {
			// 最多能购买几个i大礼包
			if quo > needs[o]/specialAviable[i][o] {
				quo = needs[o] / specialAviable[i][o] // 去最小值 也就是能购买此礼包最大个数
			}
		}

		// 计算购买该大礼包和补充购买各产品的价格
		tmpPrice := specialAviable[i][n+1]
		for o := 0; o < n; o++ {
			tmpPrice = tmpPrice + (needs[o]-(specialAviable[i][o]*quo))*price[o]
		}
		if minPrice > tmpPrice {
			minPrice = tmpPrice
		}

		// 计算购买该大礼包和其他大礼包组合方式的产品价格
		

	}

	return 0
}

package main

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price) // 商品数量和needs长度
	// 计算每个货物单独购买的总价格 作为初始最低价格
	minPrice := 0
	for i := 0; i < n; i++ {
		minPrice = minPrice + price[i]*needs[i]
	}
	// 价格为0 直接返回
	if minPrice == 0 {
		return minPrice
	}

	// 可选包为0 直接返回
	specialLength := len(special)
	if specialLength == 0 {
		return minPrice
	}

	var specialAviable [][]int

OUT:
	for i := 0; i < specialLength; i++ {
		p := 0
		for j := 0; j < n; j++ {
			if special[i][j] > needs[j] {
				// break // 去掉大礼包中有比needs中对应货物多的
				continue OUT
			}
			p = p + price[j]*special[i][j]

		}
		// 去掉比原价贵的礼包
		if special[i][n] > p {
			continue
		}
		specialAviable = append(specialAviable, special[i]) // 可用大礼包
	}

	// 计算购买每个大礼包和剩余needs加起来最低价格
	for i, j := 0, len(specialAviable); i < j; i++ {
		var nxt_needs []int
		for k := 0; k < n; k++ {
			nxt_needs = append(nxt_needs, needs[k]-specialAviable[i][k])
		}

		// 计算购买该大礼包和补充购买各产品的最低价格
		tmpPrice := specialAviable[i][n] + shoppingOffers(price, specialAviable, nxt_needs)

		if minPrice > tmpPrice {
			minPrice = tmpPrice
		}
	}

	return minPrice
}

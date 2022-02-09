package main

func maxProfit(prices []int) int {
	var maxProfit int
	buyPrice := 100000
	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
		}
		profit := price - buyPrice
		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}

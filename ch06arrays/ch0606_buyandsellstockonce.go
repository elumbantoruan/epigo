package ch06arrays

import "math"

func BuyAndSellStockOnce(prices []int) int {
	potentialProfit := 0
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		potentialProfit = prices[i] - minPrice
		maxProfit = int(math.Max(float64(maxProfit), float64(potentialProfit)))
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
	}
	return maxProfit
}

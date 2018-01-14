package ch06arrays

import "math"

// BuyAndSellStockTwice includes the second buy that must be mande
// on another date after the first sale
// given input prices : { 12, 11, 13, 9, 12, 8, 14, 13, 15 }
// first sell max profit is 7 (from 15 - 8 }
//
func BuyAndSellStockTwice(prices []int) int {
	var firstBuyProfits []int
	potentialProfit := 0
	minPrice := prices[0]
	maxPrice := 0
	maxProfit := 0
	// forward phase and record a max profit
	for i := 0; i < len(prices); i++ {
		potentialProfit = prices[i] - minPrice
		maxProfit = int(math.Max(float64(maxProfit), float64(potentialProfit)))
		firstBuyProfits = append(firstBuyProfits, maxProfit)
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
	}

	for i := len(prices) - 1; i > 0; i-- {
		maxPrice = int(math.Max(float64(maxPrice), float64(prices[i])))
		// calculate the potential profit by getting the difference between maxPrice and current index price plus the profit from the first buy
		// do firstBuyProfits[i - 1] so it returns a previous profit from current price
		potentialProfit = maxPrice - prices[i] + firstBuyProfits[i-1]
		maxProfit = int(math.Max(float64(maxProfit), float64(potentialProfit)))
	}
	return maxProfit

}

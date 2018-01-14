package ch06arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuyAndSellStockOnce(t *testing.T) {
	prices := []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}
	maxProfit := BuyAndSellStockOnce(prices)

	expected := 30
	assert.Equal(t, expected, maxProfit)

}

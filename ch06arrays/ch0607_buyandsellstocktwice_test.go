package ch06arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuyAndSellStockTwice(t *testing.T) {
	prices := []int { 12, 11, 13, 9, 12, 8, 14, 13, 15 }
	maxProfit := BuyAndSellStockTwice(prices)
	expected := 10
	assert.Equal(t, expected, maxProfit)
}

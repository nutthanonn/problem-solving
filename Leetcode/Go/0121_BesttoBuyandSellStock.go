package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	maxProfit := 0
	minNumber := prices[0]
	for _, price := range prices {
		minNumber = int(math.Min(float64(minNumber), float64(price)))
		maxProfit = int(math.Max(float64(maxProfit), float64(price-minNumber)))
	}
	return maxProfit
}

func main() {
	prices := []int{1000, 1000, 120, 42, 10, 5, 30}
	fmt.Println(maxProfit(prices))
}

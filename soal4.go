package main

import "fmt"

func bestBuyPrice(prices []int) int {
	if len(prices) < 2 {
		return -1
	}
	minPrice := prices[0]
	maxProfit := 0
	bestBuyPrice := minPrice
	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		profit := currentPrice - minPrice
		if profit > maxProfit {
			maxProfit = profit
			bestBuyPrice = minPrice
		}
		if currentPrice < minPrice {
			minPrice = currentPrice
		}
	}

	return bestBuyPrice
}

func soal4() {
	fmt.Println("SOAL 4")
	prices1 := []int{7, 8, 3, 10, 8}
	prices2 := []int{5, 12, 11, 12, 10}
	prices3 := []int{7, 18, 27, 10, 29}
	prices4 := []int{20, 17, 15, 14, 10}

	fmt.Println(prices1,bestBuyPrice(prices1))
	fmt.Println(prices2,bestBuyPrice(prices2)) 
	fmt.Println(prices3,bestBuyPrice(prices3)) 
	fmt.Println(prices4,bestBuyPrice(prices4))
}

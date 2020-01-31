package main

import "fmt"

func main() {
	in := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(in))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minBuy := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - minBuy
		if curProfit > maxProfit {
			maxProfit = curProfit
		}
		if prices[i] < minBuy {
			minBuy = prices[i]
		}
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	
	buy := prices[0]
	profit := 0

	if len(prices) < 2 {
		return profit
	}
	
	for _, price := range prices {
		if price > buy {
			profit += price - buy
		}
		buy = price
	}
	return profit
}
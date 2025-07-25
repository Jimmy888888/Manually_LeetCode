class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2 :
            return 0

        buy = prices[0]
        profit = 0

        for price in prices :
            if price > buy:
                profit += price - buy
            buy = price

        return profit
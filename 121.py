class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy = prices[0]
        max_profit = 0
        for s in range(1, len(prices)):
            if buy > prices[s]:
                buy = prices[s]
            elif prices[s] - buy > max_profit:
                max_profit = prices[s] - buy
        return max_profit

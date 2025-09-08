
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        # make dp, 0 <= amount <= 10^4
        maxN = amount + 1
        dp = [(amount+1)]*(amount+1)

        dp[0] = 0

        for i in range(1, amount+1):
            for c in coins:
                if i - c >= 0 :
                    dp[i] = min(dp[i], dp[i-c] + 1)

        if dp[amount] == maxN:
            return -1

        return dp[amount]
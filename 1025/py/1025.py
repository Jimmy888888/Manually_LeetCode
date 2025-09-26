# dp

class Solution:
    def divisorGame(self, n: int) -> bool:
        if n == 1:
            return False

        # make dp list
        dp = [False for _ in range(n+1)]

        # dp[1] = False

        # go through, Alice go first
        for i in range(2, n+1):
            for j in range(1, i):
                if i % j == 0:
                    if dp[i-j] == False:
                        dp[i] = True
                        break

        return dp[n]
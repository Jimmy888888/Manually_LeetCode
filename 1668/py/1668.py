class Solution:
    def maxRepeating(self, sequence: str, word: str) -> int:
        sl = len(sequence)
        wl = len(word)
        # make dp
        dp = [0 for _ in range(sl+1)]

        # record
        max = 0

        for i in range(1, sl+1):
            if i >= wl :
                if sequence[(i-wl):i] == word :
                    dp[i] = dp[i-wl] + 1
                    if max < dp[i]:
                        max = dp[i]
        return max
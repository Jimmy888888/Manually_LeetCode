class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        # make dict
        mapD = {}
        for i in range(0, len(wordDict)):
            mapD[wordDict[i]] = 1
        
        # make dp
        n = len(s)
        dp = [False for _ in range(n+1)]
        dp[0] = True

        # state transition
        for i in range(1, n+1):
            for j in range(0, i):
                if dp[j] == True:
                    sub = s[j:i]
                    if mapD.get(sub):
                        dp[i] = True
                        break
        
        return dp[n]
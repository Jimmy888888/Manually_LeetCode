class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        n = len(words)
        # make dp
        dp = [1 for _ in range(n)]
        # make recorder
        preWord = [-1 for _ in range(n)]

        # state transition
        for i in range(1, n) :
            for j in range(0, i) :
                if groups[i] != groups[j] :
                    if dp[i] < dp[j]+1 :
                        dp[i] = dp[j]+1
                        preWord[i] = j
        
        # get the longest
        maxLen = 0
        endId = -1
        for i in range(n) :
            if maxLen < dp[i] :
                maxLen = dp[i]
                endId = i
        
        # form the result
        result = ["" for _ in range(maxLen)]

        for i in range(maxLen-1, -1, -1) :
            result[i] = words[endId]
            endId = preWord[endId]

        return result
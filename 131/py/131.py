class Solution:
    def partition(self, s: str) -> List[List[str]]:
        # make dp
        n = len(s)
        dp = [[0 for _ in range(n)] for _ in range(n)]

        # assume : dp[i][j] means s[i~j] is a palindrome or not
        for i in range(n-1, -1, -1):
            for j in range(i, n):
                # edge cases: 
                # 1. i = n-1 -> because of short-circuit evaluation
                # j-i < 2 will == true frist, so dp[i+1][j-1] won't execute
                # 2. j = 0 -> j-1 < 2 will == true first, so dp[i+1][j-1] won't execute
                dp[i][j] = (s[i] == s[j]) and (j-i < 2 or dp[i+1][j-1])

        
        # backTrack
        path = []
        res = []
        def backTrack(startId: int) :
            if startId >= n :
                res.append(path.copy())
                return
            
            for i in range(startId, n):
                if dp[startId][i] == True:
                    subS = s[startId:i+1]
                    path.append(subS)
                    backTrack(i+1)

                    path.pop()

        
        backTrack(0)

        return res

            
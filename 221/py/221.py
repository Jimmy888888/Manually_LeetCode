class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        m = len(matrix)
        n = len(matrix[0])
        dp = [[ 0 for _ in range(n)] for _ in range(m)]
        maxSqu = 0

        for i in range(m):
            for j in range(n):
                if matrix[i][j] == "1":
                    if i == 0 or j == 0 :
                        dp[i][j] = 1
                    else :
                        dp[i][j] =  1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
                    
                    if maxSqu < dp[i][j] :
                        maxSqu = dp[i][j]

        return maxSqu * maxSqu

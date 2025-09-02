class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        m = len(obstacleGrid)
        n = len(obstacleGrid[0])

        # make DP
        DP = [[0 for _ in range(n)] for _ in range(m)]

        # base cases
        for i in range(m) :
            if obstacleGrid[i][0] == 1 :
                break
            DP[i][0] = 1

        for j in range(n) :
            if obstacleGrid[0][j] == 1 :
                break
            DP[0][j] = 1

        # transition equation
        for i in range(1, m) :
            for j in range(1, n) :
                if obstacleGrid[i][j] == 1 :
                    continue
                DP[i][j] = DP[i-1][j] + DP[i][j-1]

        return DP[m-1][n-1]

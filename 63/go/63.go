//DP
//State definition:
// go right or go down only
// DP[i][j] = the number of ways to reach grid[i][j]
// if grid[i][j] == 1, DP[i][j] = 0
//State transition equation:
// DP[i][j] = DP[i-1][j] + DP[i][j-1]
//base cases:
// first row and column

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    // make DP
    DP := make([][]int, m)
    for i := range DP {
        DP[i] = make([]int, n)
    }

    
    // base cases
    for i := 0; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        DP[i][0] = 1
    }

    for j := 0; j < n; j++ {
        if obstacleGrid[0][j] == 1 {
            break
        }
        DP[0][j] = 1
    }

    // State transition
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1{
                continue
            }
            DP[i][j] = DP[i-1][j] + DP[i][j-1]
        }
    }

    return DP[m-1][n-1]
}
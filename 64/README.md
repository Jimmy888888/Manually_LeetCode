## DP

### define state :
assume dp[i][j] is the min sum from top left to bottom right of grid[i][j]

### find state transition equation:
dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]

### base cases:

grid[0][0] the init point:

dp[0][0] = grid[0][0]

grid[0][j] only go from left:

dp[0][j] = dp[0][j-1] + grid[0][j]

grid[i][0] only go from top:

dp[i][0] = dp[i-1][0] + grid[i][0]
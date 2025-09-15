// assume
// dp[i][j] = the min number of operations to convert word1[0~i] to word2[0~j]
// dp[i][j] = dp[i-1][j-1]; if word1[i] == word2[j]
// else; dp[i][j] = min( dp[i-1][j] insert, dp[i][j-1] delete, dp[i-1][j-1] replace) + 1

import "math"

func minDistance(word1 string, word2 string) int {
    // check
    m := len(word1)
    n := len(word2)
    // make 2D dp
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }

    // base cases
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }

    // state transition equation
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = GetMin(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }

    return dp[m][n]
}

func GetMin (nums... int) int {
    minVal := math.MaxInt64
    for i := 0; i < len(nums); i++ {
        if minVal > nums[i] {
            minVal = nums[i]
        }
    }
    return minVal
}
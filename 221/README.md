# 221

## DP

### State Definition(assum):
dp[i][j] = largest square in i*j , and it's means the right down corner

dp[i][j] = (min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 )^2 

; if (dp[i-1][j] > 0 && dp[i][j-1] > 0 && dp[i-1][j-1] > 0) ; matrix[i][j] == 1

### base cases:

dp[i][j] = 1; if matrix[i][j] = 1

### record the largest

```go
import "math"

func maximalSquare(matrix [][]byte) int {
    // get m, n
    m := len(matrix)
    n := len(matrix[0])

    // make 2D dp
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    // base cases
    LargestSqu := 0
    for i := 0; i < m; i++{
        if matrix[i][0] == '1' {
            dp[i][0] = 1
            LargestSqu = 1
        }
    }

    for j := 0; j < n; j++{
        if matrix[0][j] == '1' {
            dp[0][j] = 1
            LargestSqu = 1
        }
    }

    // State contransion equation
    
    for i := 1; i < m; i++{
        for j := 1; j < n; j++{
            if matrix[i][j] == '1' {
                if matrix[i-1][j] == '0' || matrix[i][j-1] == '0' || matrix[i-1][j-1] == '0' {
                    dp[i][j] = 1
                }else{
                    dp[i][j] = GetCurSquare(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
                }
                
                if LargestSqu < dp[i][j] {
                    LargestSqu = dp[i][j]
                }
            }
        }
    }

    return LargestSqu
}

func GetCurSquare(left int, top int, leftTop int) int {

    minSqu := GetMin(left, top, leftTop)

    minL := math.Sqrt(float64(minSqu)) // float64

    curSqu := (minL + 1) * (minL + 1) // float64

    return int(curSqu)
}

func GetMin(nums... int) int {
    if len(nums) == 0 {
        return math.MaxInt64
    }

    minN := nums[0]

    for i := 0; i < len(nums); i++ {
        if nums[i] < minN {
            minN = nums[i]
        }
    }

    return minN
}
```
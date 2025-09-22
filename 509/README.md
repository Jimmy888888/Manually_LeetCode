# 509

DP

fibonacci sequence

```go
func fib(n int) int {
    if n == 0 {
        return 0
    }

    if n == 1 || n == 2{
        return 1
    }

    // make dp
    dp := make([]int, n+1)
    // base cases
    dp[0] = 0
    dp[1] = 1

    //State transition equation
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}
```

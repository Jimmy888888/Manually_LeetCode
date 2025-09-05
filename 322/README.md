# 322

## way1
assume dp[i] = the fewest number of coins that can form the 

amount i, if can't it will be -1

because can't make sure previous coin number, so assume dp[j]

dp[i] = min(dp[j]) + 1, if((i-j) in coins), dp[j] > -1, j = 0~i-1

### result

work , but too slow, get Time Limit Exceeded

```go
func coinChange(coins []int, amount int) int {
    // check 
    if amount == 0 {
        return 0;
    }
    // make map
    mapC := make(map[int]int)
    for i := 0; i < len(coins); i++{
        mapC[coins[i]] = 1
    }

    // make dp
    dp := make([]int, amount+1)
    for i := range dp{
        dp[i] = -1
    }

    dp[0] = 0

    // state transition equation
    for i := 1; i < amount+1; i++{
        for j := 0; j < i; j++{
            _, find := mapC[i-j]
            if dp[j] > -1 && find{
                if dp[i] > -1 {
                    if dp[i] > dp[j] + 1 {
                        dp[i] = dp[j] + 1
                    }
                }else{
                    dp[i] = dp[j] + 1
                }
            }
        }
    }

    return dp[amount]
}
```

## way2

assume dp[i] = the fewest number of coins that can form the 

amount i, if can't it will be -1

because can't make sure previous coin number, if dp[i] valid

exist dp[j] that (i-j) in coins, range of i-j will be one of 
coin in coins

so, i-j can be replace to coins

--> dp[i] = min(dp[i-coin] + 1) , coin from coins

```go
func coinChange(coins []int, amount int) int {
    // check 
    if amount == 0 {
        return 0;
    }

    // make dp
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++{
        dp[i] = -1
    }

    // state transition equation
    for i := 1; i < amount+1; i++{
        for _, c := range coins{
            if i >= c {
                if dp[i-c] > -1 {
                    if dp[i] > -1{
                        dp[i] = min(dp[i], 1+dp[i-c])
                    }else{
                        dp[i] = 1+dp[i-c]
                    }
                }
                
            }
        }
    }

    return dp[amount]
}

func min(a int, b int) int{
    if a > b {
        return b
    }
    return a
}
```

### result

pass
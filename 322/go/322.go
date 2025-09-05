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
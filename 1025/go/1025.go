func divisorGame(n int) bool {
    if n == 1 {
        return false
    }

    // make dp slices
    dp := make([]bool, n+1)

    // dp[1] = false

    // Alice go first
    for i := 2; i < n+1; i++{
        for j := 1; j < i; j++{
            if i % j == 0{
                if dp[i-j] == false{
                    dp[i] = true
                    break
                }
            }
        }
    }
    
    return dp[n]
}
func maxRepeating(sequence string, word string) int {
    sl := len(sequence)
    wl := len(word)

    // make dp
    dp := make([]int, sl+1)

    max := 0

    for i := 1; i < sl+1; i++ {
        if i >= wl {
            if sequence[i-wl:i] == word {
                dp[i] = dp[i-wl] + 1
                if max < dp[i]{
                    max = dp[i]
                }
            }
        }
    }

    return max
}
func getLongestSubsequence(words []string, groups []int) []string {
    n := len(words)
    // dp, preWord
    dp := make([]int, n)
    pre := make([]int, n)

    for i := 0; i < n; i++{
        dp[i] = 1
    }

    // state transition
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++{
            if groups[i] != groups[j] {
                if dp[i] < dp[j] + 1 {
                    dp[i] = dp[j] + 1
                    pre[i] = j
                }
            }
        }
    }

    maxL := -1
    endId := 0
    for i := 0; i < n; i++ {
        if maxL < dp[i] {
            maxL = dp[i]
            endId = i
        }
    }

    result := make([]string, maxL)
    for i := maxL-1; i >= 0; i-- {
        result[i] = words[endId]
        endId = pre[endId]
    }

    return result
}
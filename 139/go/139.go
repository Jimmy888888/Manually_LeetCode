// assume dp[i] = if s[i] can be formed by wordDict
// because can't make sure the lrngth of previous one, so assume it dp[j]
// dp[i] = (s[i-j] in wordDict) && dp[j]
// len(dp) = len(s)+1 , set dp[0] = true

func wordBreak(s string, wordDict []string) bool {

    // make a map of wordDict
    mapD := make(map[string]int)

    for i := 0; i < len(wordDict); i++{
        mapD[wordDict[i]] = 1
    }

    // make dp
    n := len(s)
    dp := make([]bool, n+1)

    dp[0] = true

    for i := 1; i <= n; i++{
        for j := 0; j < i; j++{
            if dp[j] == true {
                sub := s[j:i]
                if _, find := mapD[sub]; find{
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[n]
}
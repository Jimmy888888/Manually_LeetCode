// DP

// assume
// dp[i] is the possible ways of s[0~i]
// dp[i] =>
// 1. s[i-1 ~ i] is in ("1"~"26")  &&  s[i] is in ("1"~"26") => dp[i] = dp[i-1] + dp[i-2]
// 2. only s[i-1 ~ i] is in ("1"~"26") => dp[i] = dp[i-2]
// 3. only s[i] is in ("1"~"26") => dp[i] = dp[i-1]

import "strconv"

func numDecodings(s string) int {
    // check
    n := len(s)

    if s[0] == '0' {
        return 0
    }
    if n == 1 {
        return 1
    }

    // make dp
    dp := make([]int, n)
    // make checkM
    checkM := make(map[string]int)
    for i := 1; i <= 26; i++{
        checkM[strconv.Itoa(i)] = 1
    }

    // base cases
    dp[0] = 1
    _, find1 := checkM[s[1:2]]
    _, find2 := checkM[s[0:2]]
    if find1 && find2 {
        dp[1] = 2
    }else if find1 || find2{
        dp[1] = 1
    }

    if dp[1] == 0 {
        return 0
    }

    for i := 2; i < n; i++{
        _, find1 := checkM[s[i:i+1]]
        _, find2 := checkM[s[i-1:i+1]]
        if find1 && find2 {
            dp[i] = dp[i-1] + dp[i-2]
        }else if find1{
            dp[i] = dp[i-1]
        }else if find2{
            dp[i] = dp[i-2]
        }

        if dp[i] == 0 {
            return 0
        }
    }

    return dp[n-1]

}
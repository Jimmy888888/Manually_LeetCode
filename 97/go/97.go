// DP 
//State definition
// build a 2D bool{(s1.length+1) * (s2.length+1)} array DP[i][j], init with 0
// DP[i][j] = whether first i chrs in s1 and first j chrs in s2 can build first i+j chrs in s3

//State Equation
// DP[i][j] = (DP[i-1][j] && s1[i-1] === s3[i+j-1]) || (DP[i][j-1] && s2[j-1] === s3[i+j-1])
// can simplify to 1D slice 


//Base cases
// DP[i][0] = 1
// DP[0][j] = 1

// finally check DP[s1.length][s2.length]

func isInterleave(s1 string, s2 string, s3 string) bool {
    // check len
    n1 := len(s1)
    n2 := len(s2)
    n3 := len(s3)
    if n1 + n2 != n3 {
        return false
    }

    // make dp[i] -> ref s1
    dp := make([]bool, n1+1)

    // base cases
    dp[0] = true
    for i := 1; i < n1+1; i++ {
        if s1[i-1] != s3[i-1]{
            break
        }
        dp[i] = true
    }

    // state transition equation
    for i := 1; i < n2+1; i++ {
        chr2 := s2[i-1]
        dp[0] = s2[i-1] == s3[i-1] && dp[0]
        for j := 1; j < n1+1; j++{
            chr3 := s3[i+j-1]
            chr1 := s1[j-1]
            dp[j] = ((chr1 == chr3) && dp[j-1]) || ((chr2 == chr3) && dp[j])
        }
    }

    return dp[n1]
}
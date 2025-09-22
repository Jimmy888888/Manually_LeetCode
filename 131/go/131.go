// DP
// all possible palindrome
// backtrack & DP
// 

func partition(s string) [][]string {
    n := len(s)
    var path []string
    var res [][]string

    // make dp
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
    }

    for i := n-1; i > -1; i-- {
        for j := i; j < n; j++{
            if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]){
                dp[i][j] = true
            }
        }
    }

    var backTrack func(startInt int)

    backTrack = func(startInt int){
        if startInt >= n {
            temp := make([]string, len(path))
            copy(temp, path)
            res = append(res, temp)
            return
        }

        for i := startInt; i < n; i++ {
            if dp[startInt][i] {
                subStr := s[startInt: i+1]
                path = append(path, subStr)

                backTrack(i+1)
                path = path[:len(path)-1]
            }
        }
    }

    backTrack(0)

    return res
}
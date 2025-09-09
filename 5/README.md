# 5

## State Definition

how to check a palindromic:

-> two flag(f & b) go through forward and backward, and the two flag must be the same always, 

-> until f >= b

assume dp[l][r] means s[l...r] is a palindromic

-> dp[l][r] = ( s[l] === s[r] ) && dp[l]
//
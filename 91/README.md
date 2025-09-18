# 91
## DP


### define state:
- "1" ~ "26"

- can't start with "0"

- assume:

len(dp) = len(s)

dp[i] = s[0~i] possiable decode ways

### state transition equation:

dp[i] =

1. dp[i-1] -> s[i] in ("1"~"9")

2. dp[i-2] -> s[i-1] + s[i] in ("10"~"26")

If the entire string cannot be decoded in any valid way, return 0

-> if 1. 2. both not possiable -> return 0, 

### base cases:

- dp[0] = 0

- if s[0] != "0", dp[1] = 1

### return:

- dp[len(s)]

### code:
```py
class Solution:
    def numDecodings(self, s: str) -> int:
        # check
        n = len(s)
        if s[0] == "0" :
            return 0
        if n < 2 :
            return 1

        # make dp
        dp = [0 for _ in range(n)]

        # make map for "1"~"26"
        codeMap = dict()
        for i in range(1,27):
            codeMap[str(i)] = 1

        # base cases
        dp[0] = 1

        if codeMap.get(s[0] + s[1]) is not None and codeMap.get(s[1]) is not None:
            dp[1] = 2
        elif codeMap.get(s[1]) is not None:
            dp[1] = 1
        elif codeMap.get(s[0] + s[1]) is not None:
            dp[1] = 1
        
        if dp[1] == 0 :
            return 0
        
        # state transition equation
        for i in range(2, n):
            s1 = s[i]
            s2 = s[i-1] + s1

            if codeMap.get(s1) is not None and codeMap.get(s2) is not None:
                dp[i] = dp[i-1] + dp[i-2]
            elif codeMap.get(s1) is not None:
                dp[i] = dp[i-1]
            elif codeMap.get(s2) is not None:
                dp[i] = dp[i-2]
            
            if dp[i] == 0 :
                return 0

        return dp[n-1]

#dp[1] = 1, dp[2] = 2, dp[3] = 3, dp[4] = 5 2
```
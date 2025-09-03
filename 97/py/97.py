class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        # check len
        n1 = len(s1)
        n2 = len(s2)
        n3 = len(s3)

        if n1+n2 != n3 :
            return False

        # make check dp
        dp = [False for _ in range(n1+1)]

        # init dp
        dp[0] = True
        for i in range(1, n1+1) :
            if s1[i-1] != s3[i-1]:
                break
            dp[i] = True
            
        for j in range(1, n2+1):
            c2 = s2[j-1]
            dp[0] = (c2 == s3[j-1]) and dp[0]
            for i in range(1, n1+1):
                dp[i] = (c2 == s3[i+j-1] and dp[i]) or (s1[i-1] == s3[i+j-1] and dp[i-1])

        return dp[n1]
class Solution:
    def longestPalindrome(self, s: str) -> str:

        n = len(s)
        if n == 1 :
            return s
        
        maxS = ""
        curSOdd = ""
        curSEven = ""

        for i in range(n):
            curSOdd = self.centerExpend(i, i, s)

            if i < n-1 :
                curSEven = self.centerExpend(i, i+1, s)

            if len(maxS) < len(curSOdd) :
                maxS = curSOdd
            if len(maxS) < len(curSEven) :
                maxS = curSEven
            
        return maxS
        

    def centerExpend(self, l: int, r: int, s: str) -> str:
        subS = ""

        while l >= 0 and r < len(s) :
            if s[l] == s[r] :
                subS = s[l:r+1]
            else :
                break
            l -= 1
            r += 1
        
        return subS
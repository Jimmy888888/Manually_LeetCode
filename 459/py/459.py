class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        # start at s[0], end at s[1~n/2]
        # n % subLen == 0

        n = len(s)
        for i in range(0, n//2 + 1) :
            if n % (i+1) == 0 :
                if self.checkSub(0, i, s, n) :
                    return True

        return False

    def checkSub(self, start: int, end: int, s: str, sl: int) -> bool:
        subLen = end - start + 1

        for i in range(end+1, sl) :
            if s[i % subLen] != s[i] :
                break
            if i == sl-1 :
                return True
        
        return False
        
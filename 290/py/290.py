class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        n = len(pattern)

        arrS = s.split(" ")
        m = len(arrS)

        if n != m :
            return False

        # make dict
        mapS = {}
        mapP = {}

        for i in range(0, n):
            st = arrS[i]
            p = pattern[i]

            if mapS.get(st) :
                if mapS.get(st) != p:
                    return False
            else:
                mapS[st] = p
            
            if mapP.get(p) :
                if mapP.get(p) != st:
                    return False
            else:
                mapP[p] = st
        
        return True


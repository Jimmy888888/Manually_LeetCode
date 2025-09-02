class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        dictS = {}
        dictT = {}

        n = len(s)

        for i in range(n) :
            chrS = s[i]
            chrT = t[i]

            if dictS.get(chrS):
                if dictS[chrS] != chrT :
                    return False
            else :
                dictS[chrS] = chrT

            if dictT.get(chrT):
                if dictT[chrT] != chrS :
                    return False
            else :
                dictT[chrT] = chrS

        return True
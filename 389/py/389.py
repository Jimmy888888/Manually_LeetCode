class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        sCount = {}
        tCount = {}

        sl = len(s)
        tl = sl + 1

        for c in s :
            if sCount.get(c) != None :
                sCount[c] = sCount[c] + 1
            else :
                sCount[c] = 1
        
        for c in t :
            if tCount.get(c) != None :
                tCount[c] = tCount[c] + 1
            else :
                tCount[c] = 1

        for tKey, tVal in tCount.items() :
            if sCount.get(tKey) != None :
                if tVal != sCount[tKey] :
                    return tKey
            else :
                return tKey


        
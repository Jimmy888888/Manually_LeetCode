class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        
        sortDict = {}
        outPut = []

        for s in strs :
            sList = list(s)
            sortList = sorted(sList)
            sortStr = "".join(sortList)

            if sortDict.get(sortStr) == None :
                Id = len(sortDict)
                sortDict[sortStr] = Id

                newGroup = [s]
                outPut.append(newGroup)
            else :
                Id = sortDict[sortStr]
                outPut[Id].append(s)

        return outPut
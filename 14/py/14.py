class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 0 :
            return ""
        
        comPreStr = ""
        end = False
        n = len(strs[0])
        m = len(strs)

        for i in range(0, n, 1):
            end = False
            for j in range(1, m, 1):
                if i > len(strs[j]) - 1 :
                    end = True
                    break
                if strs[0][i] != strs[j][i] :
                    end = True
                    break

            if end :
                break

            comPreStr += strs[0][i]
        
        return comPreStr

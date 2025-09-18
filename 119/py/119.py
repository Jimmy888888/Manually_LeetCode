class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        result = [[1]]

        for i in range(1, rowIndex+1):
            newR = []
            for j in range(0, i+1):
                if j == 0 or j == i:
                    newR.append(1)
                else:
                    newR.append(result[i-1][j] + result[i-1][j-1])

            result.append(newR)
        
        return result[rowIndex]
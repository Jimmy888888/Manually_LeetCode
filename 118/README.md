# 118

triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]

len(triangle[i]) = len(triangle[i-1]) + 2

if j == 0 or j ==len(triangle[i])-1 : triangle[i][j] = 1

```py
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        triangle = [[1]]
        for i in range(1,numRows):
            newRow = []
            for j in range(0, i+1):
                if j == 0 or j == i :
                    newRow.append(1)
                else :
                    newNum = triangle[i-1][j-1] + triangle[i-1][j]
                    newRow.append(newNum)


            triangle.append(newRow)
        
        return triangle
```

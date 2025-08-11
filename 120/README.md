## DP 

move from botton tp top

### define:

n = the layer number of the triangle, i = n-1~0 , j = i~0

f(i,j) = the min path of triangle[i][j] , from button to top


### at the last layer:

f(n-1,j) = triangle[n-1][j]


### from the second to last layer:

i = n-2~0, j = i~0

f(i,j) = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])

in the end, the triangle[0][0] will be the min path

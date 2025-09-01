
## DP

### State Definition:
assume DP[i][j] = the number of ways to reach point [i,j]

### State Transition Equation:
if grid[i][0] == 0

DP[i][j] = DP[i-1][j] + DP[i][j-1]

### Base cases:
DP[0~m][0] = 1, DP[0][0~n] = 1

if grid[i][0] == 1 -> DP[i][0], DP[i+1][0] ... = 0

if grid[0][j] == 1 -> DP[0][j], DP[0][j+1] ... = 0

if grid[i][j] == 1 -> DP[i][j] = 0


```js
/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */

var uniquePathsWithObstacles = function(obstacleGrid) {
    let m = obstacleGrid.length;
    let n = obstacleGrid[0].length;
    // make a 2D array
    let DP = Array(m).fill(0).map(() => Array(n).fill(0));

    // base cases
    for (let i = 0; i < m; i++){
        if (obstacleGrid[i][0] === 1){
            break;
        }
        DP[i][0] = 1;
    }

    for (let j = 0; j < n; j++){
        if (obstacleGrid[0][j] === 1){
            break;
        }
        DP[0][j] = 1;
    }

    // transition equation
    for (let i = 1; i < m; i++){
        for (let j = 1; j < n; j++){
            if(obstacleGrid[i][j] === 1){
                continue;
            }

            DP[i][j] = DP[i - 1][j] + DP[i][j - 1];
        }
    }

    return DP[m-1][n-1];
};
```

### result: Pass